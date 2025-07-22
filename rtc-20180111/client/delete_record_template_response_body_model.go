// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRecordTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteRecordTemplateResponseBody
	GetRequestId() *string
}

type DeleteRecordTemplateResponseBody struct {
	// example:
	//
	// 760bad53276431c499e30dc36f6b26be
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRecordTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRecordTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRecordTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRecordTemplateResponseBody) SetRequestId(v string) *DeleteRecordTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRecordTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
