// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitDocStructureJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SubmitDocStructureJobResponseBody
	GetCode() *string
	SetData(v *SubmitDocStructureJobResponseBodyData) *SubmitDocStructureJobResponseBody
	GetData() *SubmitDocStructureJobResponseBodyData
	SetMessage(v string) *SubmitDocStructureJobResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitDocStructureJobResponseBody
	GetRequestId() *string
}

type SubmitDocStructureJobResponseBody struct {
	// example:
	//
	// noPermission
	Code *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *SubmitDocStructureJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// You are not authorized to perform this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 43A29C77-405E-4CC0-BC55-EE694AD00655
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitDocStructureJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitDocStructureJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitDocStructureJobResponseBody) GetCode() *string {
	return s.Code
}

func (s *SubmitDocStructureJobResponseBody) GetData() *SubmitDocStructureJobResponseBodyData {
	return s.Data
}

func (s *SubmitDocStructureJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitDocStructureJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitDocStructureJobResponseBody) SetCode(v string) *SubmitDocStructureJobResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitDocStructureJobResponseBody) SetData(v *SubmitDocStructureJobResponseBodyData) *SubmitDocStructureJobResponseBody {
	s.Data = v
	return s
}

func (s *SubmitDocStructureJobResponseBody) SetMessage(v string) *SubmitDocStructureJobResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitDocStructureJobResponseBody) SetRequestId(v string) *SubmitDocStructureJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitDocStructureJobResponseBody) Validate() error {
	return dara.Validate(s)
}

type SubmitDocStructureJobResponseBodyData struct {
	// example:
	//
	// docmind-20220816-15bc7965
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s SubmitDocStructureJobResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SubmitDocStructureJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitDocStructureJobResponseBodyData) GetId() *string {
	return s.Id
}

func (s *SubmitDocStructureJobResponseBodyData) SetId(v string) *SubmitDocStructureJobResponseBodyData {
	s.Id = &v
	return s
}

func (s *SubmitDocStructureJobResponseBodyData) Validate() error {
	return dara.Validate(s)
}
