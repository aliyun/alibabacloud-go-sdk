// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAppRecordTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyAppRecordTemplateResponseBody
	GetRequestId() *string
}

type ModifyAppRecordTemplateResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAppRecordTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppRecordTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAppRecordTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAppRecordTemplateResponseBody) SetRequestId(v string) *ModifyAppRecordTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAppRecordTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
