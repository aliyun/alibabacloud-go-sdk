// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloneServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLabels(v map[string]*string) *CloneServiceRequest
	GetLabels() map[string]*string
	SetBody(v string) *CloneServiceRequest
	GetBody() *string
}

type CloneServiceRequest struct {
	// The label of the service to be cloned.
	Labels map[string]*string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// The request body. For more information, see [CreateService](https://help.aliyun.com/document_detail/412086.html).
	//
	// example:
	//
	// {   "name": "foo",   "model_path": "http://path/to/model.tar.gz",   "processor": "tensorflow_cpu",   "metadata": {     "instance": 2,     "memory": 7000,     "cpu": 4   } }
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloneServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s CloneServiceRequest) GoString() string {
	return s.String()
}

func (s *CloneServiceRequest) GetLabels() map[string]*string {
	return s.Labels
}

func (s *CloneServiceRequest) GetBody() *string {
	return s.Body
}

func (s *CloneServiceRequest) SetLabels(v map[string]*string) *CloneServiceRequest {
	s.Labels = v
	return s
}

func (s *CloneServiceRequest) SetBody(v string) *CloneServiceRequest {
	s.Body = &v
	return s
}

func (s *CloneServiceRequest) Validate() error {
	return dara.Validate(s)
}
