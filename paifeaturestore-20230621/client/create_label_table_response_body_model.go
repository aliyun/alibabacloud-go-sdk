// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLabelTableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLabelTableId(v string) *CreateLabelTableResponseBody
	GetLabelTableId() *string
	SetRequestId(v string) *CreateLabelTableResponseBody
	GetRequestId() *string
}

type CreateLabelTableResponseBody struct {
	// example:
	//
	// 1
	LabelTableId *string `json:"LabelTableId,omitempty" xml:"LabelTableId,omitempty"`
	// example:
	//
	// 0FA90B3B-F30A-5C9D-A9FD-8114F8868062
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateLabelTableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateLabelTableResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLabelTableResponseBody) GetLabelTableId() *string {
	return s.LabelTableId
}

func (s *CreateLabelTableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateLabelTableResponseBody) SetLabelTableId(v string) *CreateLabelTableResponseBody {
	s.LabelTableId = &v
	return s
}

func (s *CreateLabelTableResponseBody) SetRequestId(v string) *CreateLabelTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLabelTableResponseBody) Validate() error {
	return dara.Validate(s)
}
