// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloneServiceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLabelsShrink(v string) *CloneServiceShrinkRequest
	GetLabelsShrink() *string
	SetBody(v string) *CloneServiceShrinkRequest
	GetBody() *string
}

type CloneServiceShrinkRequest struct {
	// The label of the service to be cloned.
	LabelsShrink *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// The request body. For more information, see [CreateService](https://help.aliyun.com/document_detail/412086.html).
	//
	// example:
	//
	// {   "name": "foo",   "model_path": "http://path/to/model.tar.gz",   "processor": "tensorflow_cpu",   "metadata": {     "instance": 2,     "memory": 7000,     "cpu": 4   } }
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloneServiceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CloneServiceShrinkRequest) GoString() string {
	return s.String()
}

func (s *CloneServiceShrinkRequest) GetLabelsShrink() *string {
	return s.LabelsShrink
}

func (s *CloneServiceShrinkRequest) GetBody() *string {
	return s.Body
}

func (s *CloneServiceShrinkRequest) SetLabelsShrink(v string) *CloneServiceShrinkRequest {
	s.LabelsShrink = &v
	return s
}

func (s *CloneServiceShrinkRequest) SetBody(v string) *CloneServiceShrinkRequest {
	s.Body = &v
	return s
}

func (s *CloneServiceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
