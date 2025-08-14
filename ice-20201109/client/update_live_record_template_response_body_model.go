// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLiveRecordTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateLiveRecordTemplateResponseBody
	GetRequestId() *string
}

type UpdateLiveRecordTemplateResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0F3D5C03-4B6E-5F40-B7F6-B1956776E7D3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateLiveRecordTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveRecordTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLiveRecordTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateLiveRecordTemplateResponseBody) SetRequestId(v string) *UpdateLiveRecordTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLiveRecordTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
