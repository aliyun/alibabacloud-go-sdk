// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUpdateType(v string) *UpdateServiceRequest
	GetUpdateType() *string
	SetBody(v string) *UpdateServiceRequest
	GetBody() *string
}

type UpdateServiceRequest struct {
	// The type of the service update. Valid values: merge and replace. By default, merge is used if you do not specify this parameter.
	//
	// 	- merge: If the JSON string configured for the existing service is `{"a":"b"}` and the JSON string specified in the body parameter is `{"c":"d"}`, the JSON string is `{"a":"b","c":"d"}` after the service update.
	//
	// 	- replace: If the JSON string configured for the existing service is `{"a":"b"}` and the JSON string specified in the body parameter is `{"c":"d"}`, the JSON string is `{"c":"d"}` after the service update.
	//
	// example:
	//
	// merge
	UpdateType *string `json:"UpdateType,omitempty" xml:"UpdateType,omitempty"`
	// The request body. The body includes the request parameters that you want to update. For more information about the request parameters, see [CreateService](https://help.aliyun.com/document_detail/412086.html).
	//
	// example:
	//
	// {   "name": "foo",   "model_path": "http://path/to/model.tar.gz",   "processor": "tensorflow_cpu",   "metadata": {     "instance": 2,     "memory": 7000,     "cpu": 4   } }
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceRequest) GetUpdateType() *string {
	return s.UpdateType
}

func (s *UpdateServiceRequest) GetBody() *string {
	return s.Body
}

func (s *UpdateServiceRequest) SetUpdateType(v string) *UpdateServiceRequest {
	s.UpdateType = &v
	return s
}

func (s *UpdateServiceRequest) SetBody(v string) *UpdateServiceRequest {
	s.Body = &v
	return s
}

func (s *UpdateServiceRequest) Validate() error {
	return dara.Validate(s)
}
