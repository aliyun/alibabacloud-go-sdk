// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitDigitalDocStructureJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SubmitDigitalDocStructureJobResponseBody
	GetCode() *string
	SetData(v interface{}) *SubmitDigitalDocStructureJobResponseBody
	GetData() interface{}
	SetId(v string) *SubmitDigitalDocStructureJobResponseBody
	GetId() *string
	SetMessage(v string) *SubmitDigitalDocStructureJobResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitDigitalDocStructureJobResponseBody
	GetRequestId() *string
	SetStatus(v string) *SubmitDigitalDocStructureJobResponseBody
	GetStatus() *string
}

type SubmitDigitalDocStructureJobResponseBody struct {
	// example:
	//
	// noPermission
	Code *string     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Id   *string     `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// You are not authorized to perform this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s SubmitDigitalDocStructureJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitDigitalDocStructureJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitDigitalDocStructureJobResponseBody) GetCode() *string {
	return s.Code
}

func (s *SubmitDigitalDocStructureJobResponseBody) GetData() interface{} {
	return s.Data
}

func (s *SubmitDigitalDocStructureJobResponseBody) GetId() *string {
	return s.Id
}

func (s *SubmitDigitalDocStructureJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitDigitalDocStructureJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitDigitalDocStructureJobResponseBody) GetStatus() *string {
	return s.Status
}

func (s *SubmitDigitalDocStructureJobResponseBody) SetCode(v string) *SubmitDigitalDocStructureJobResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitDigitalDocStructureJobResponseBody) SetData(v interface{}) *SubmitDigitalDocStructureJobResponseBody {
	s.Data = v
	return s
}

func (s *SubmitDigitalDocStructureJobResponseBody) SetId(v string) *SubmitDigitalDocStructureJobResponseBody {
	s.Id = &v
	return s
}

func (s *SubmitDigitalDocStructureJobResponseBody) SetMessage(v string) *SubmitDigitalDocStructureJobResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitDigitalDocStructureJobResponseBody) SetRequestId(v string) *SubmitDigitalDocStructureJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitDigitalDocStructureJobResponseBody) SetStatus(v string) *SubmitDigitalDocStructureJobResponseBody {
	s.Status = &v
	return s
}

func (s *SubmitDigitalDocStructureJobResponseBody) Validate() error {
	return dara.Validate(s)
}
