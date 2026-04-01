// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGadInstanceMemberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateGadInstanceMemberResponseBody
	GetRequestId() *string
	SetResult(v *CreateGadInstanceMemberResponseBodyResult) *CreateGadInstanceMemberResponseBody
	GetResult() *CreateGadInstanceMemberResponseBodyResult
}

type CreateGadInstanceMemberResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 16C62438-491B-5C02-9B49-BA924A1372A2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// An array that consists of the information returned.
	Result *CreateGadInstanceMemberResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CreateGadInstanceMemberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateGadInstanceMemberResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGadInstanceMemberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateGadInstanceMemberResponseBody) GetResult() *CreateGadInstanceMemberResponseBodyResult {
	return s.Result
}

func (s *CreateGadInstanceMemberResponseBody) SetRequestId(v string) *CreateGadInstanceMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGadInstanceMemberResponseBody) SetResult(v *CreateGadInstanceMemberResponseBodyResult) *CreateGadInstanceMemberResponseBody {
	s.Result = v
	return s
}

func (s *CreateGadInstanceMemberResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateGadInstanceMemberResponseBodyResult struct {
	// The number of unit nodes that are created by calling this operation.
	//
	// example:
	//
	// 2
	CreateCount *string `json:"CreateCount,omitempty" xml:"CreateCount,omitempty"`
	// The ID of the global active database cluster.
	//
	// example:
	//
	// gad-rm-bp1npi2j8********
	GadInstanceName *string `json:"GadInstanceName,omitempty" xml:"GadInstanceName,omitempty"`
}

func (s CreateGadInstanceMemberResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s CreateGadInstanceMemberResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateGadInstanceMemberResponseBodyResult) GetCreateCount() *string {
	return s.CreateCount
}

func (s *CreateGadInstanceMemberResponseBodyResult) GetGadInstanceName() *string {
	return s.GadInstanceName
}

func (s *CreateGadInstanceMemberResponseBodyResult) SetCreateCount(v string) *CreateGadInstanceMemberResponseBodyResult {
	s.CreateCount = &v
	return s
}

func (s *CreateGadInstanceMemberResponseBodyResult) SetGadInstanceName(v string) *CreateGadInstanceMemberResponseBodyResult {
	s.GadInstanceName = &v
	return s
}

func (s *CreateGadInstanceMemberResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
