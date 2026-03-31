// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyDefenseTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CopyDefenseTemplateResponseBody
	GetRequestId() *string
	SetTemplateId(v int64) *CopyDefenseTemplateResponseBody
	GetTemplateId() *int64
}

type CopyDefenseTemplateResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D7861F61-5B61-46CE-A47C-6B19****5EB0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the new protection template.
	//
	// example:
	//
	// 12346
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s CopyDefenseTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CopyDefenseTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CopyDefenseTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CopyDefenseTemplateResponseBody) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *CopyDefenseTemplateResponseBody) SetRequestId(v string) *CopyDefenseTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CopyDefenseTemplateResponseBody) SetTemplateId(v int64) *CopyDefenseTemplateResponseBody {
	s.TemplateId = &v
	return s
}

func (s *CopyDefenseTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
