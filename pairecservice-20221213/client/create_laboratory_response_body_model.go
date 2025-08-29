// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLaboratoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLaboratoryId(v string) *CreateLaboratoryResponseBody
	GetLaboratoryId() *string
	SetRequestId(v string) *CreateLaboratoryResponseBody
	GetRequestId() *string
}

type CreateLaboratoryResponseBody struct {
	// example:
	//
	// 3
	LaboratoryId *string `json:"LaboratoryId,omitempty" xml:"LaboratoryId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 42391E6D-822C-58F8-9F7E-D991BB86D6AD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateLaboratoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateLaboratoryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLaboratoryResponseBody) GetLaboratoryId() *string {
	return s.LaboratoryId
}

func (s *CreateLaboratoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateLaboratoryResponseBody) SetLaboratoryId(v string) *CreateLaboratoryResponseBody {
	s.LaboratoryId = &v
	return s
}

func (s *CreateLaboratoryResponseBody) SetRequestId(v string) *CreateLaboratoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLaboratoryResponseBody) Validate() error {
	return dara.Validate(s)
}
