// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportTemplateEventResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ImportTemplateEventResponseBody
	GetRequestId() *string
	SetData(v bool) *ImportTemplateEventResponseBody
	GetData() *bool
}

type ImportTemplateEventResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returned data object
	//
	// example:
	//
	// true
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
}

func (s ImportTemplateEventResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImportTemplateEventResponseBody) GoString() string {
	return s.String()
}

func (s *ImportTemplateEventResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImportTemplateEventResponseBody) GetData() *bool {
	return s.Data
}

func (s *ImportTemplateEventResponseBody) SetRequestId(v string) *ImportTemplateEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportTemplateEventResponseBody) SetData(v bool) *ImportTemplateEventResponseBody {
	s.Data = &v
	return s
}

func (s *ImportTemplateEventResponseBody) Validate() error {
	return dara.Validate(s)
}
