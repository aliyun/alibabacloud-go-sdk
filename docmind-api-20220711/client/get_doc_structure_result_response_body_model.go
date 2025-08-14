// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocStructureResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetDocStructureResultResponseBody
	GetCode() *string
	SetCompleted(v bool) *GetDocStructureResultResponseBody
	GetCompleted() *bool
	SetData(v map[string]interface{}) *GetDocStructureResultResponseBody
	GetData() map[string]interface{}
	SetMessage(v string) *GetDocStructureResultResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetDocStructureResultResponseBody
	GetRequestId() *string
	SetStatus(v string) *GetDocStructureResultResponseBody
	GetStatus() *string
}

type GetDocStructureResultResponseBody struct {
	// example:
	//
	// noPermission
	Code      *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Completed *bool                  `json:"Completed,omitempty" xml:"Completed,omitempty"`
	Data      map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// You are not authorized to perform this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetDocStructureResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDocStructureResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetDocStructureResultResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetDocStructureResultResponseBody) GetCompleted() *bool {
	return s.Completed
}

func (s *GetDocStructureResultResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *GetDocStructureResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDocStructureResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDocStructureResultResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetDocStructureResultResponseBody) SetCode(v string) *GetDocStructureResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetDocStructureResultResponseBody) SetCompleted(v bool) *GetDocStructureResultResponseBody {
	s.Completed = &v
	return s
}

func (s *GetDocStructureResultResponseBody) SetData(v map[string]interface{}) *GetDocStructureResultResponseBody {
	s.Data = v
	return s
}

func (s *GetDocStructureResultResponseBody) SetMessage(v string) *GetDocStructureResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetDocStructureResultResponseBody) SetRequestId(v string) *GetDocStructureResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDocStructureResultResponseBody) SetStatus(v string) *GetDocStructureResultResponseBody {
	s.Status = &v
	return s
}

func (s *GetDocStructureResultResponseBody) Validate() error {
	return dara.Validate(s)
}
