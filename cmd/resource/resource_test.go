package resource

import (
	"reflect"
	"testing"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
)

func TestCreate(t *testing.T) {
	type args struct {
		req          handler.Request
		prevModel    *Model
		currentModel *Model
	}
	tests := []struct {
		name    string
		args    args
		want    handler.ProgressEvent
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Create(tt.args.req, tt.args.prevModel, tt.args.currentModel)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() = %v, want %v", got, tt.want)
			}
		})
	}
}
