// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveRecordTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteLiveRecordTemplateResponseBody
	GetRequestId() *string
}

type DeleteLiveRecordTemplateResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 3E5330CF-B4C8-5BEF-AA6B-8E70BD20FAEE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLiveRecordTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveRecordTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLiveRecordTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLiveRecordTemplateResponseBody) SetRequestId(v string) *DeleteLiveRecordTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLiveRecordTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
