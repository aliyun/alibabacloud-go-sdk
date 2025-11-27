// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGADInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateGADInstanceResponseBody
	GetRequestId() *string
	SetResult(v *CreateGADInstanceResponseBodyResult) *CreateGADInstanceResponseBody
	GetResult() *CreateGADInstanceResponseBodyResult
}

type CreateGADInstanceResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 9F8C06AD-3F37-57A0-ABBF-ABD7824F55CE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The data returned.
	Result *CreateGADInstanceResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CreateGADInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateGADInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGADInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateGADInstanceResponseBody) GetResult() *CreateGADInstanceResponseBodyResult {
	return s.Result
}

func (s *CreateGADInstanceResponseBody) SetRequestId(v string) *CreateGADInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGADInstanceResponseBody) SetResult(v *CreateGADInstanceResponseBodyResult) *CreateGADInstanceResponseBody {
	s.Result = v
	return s
}

func (s *CreateGADInstanceResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateGADInstanceResponseBodyResult struct {
	// The number of unit nodes that are created by calling this operation.
	//
	// example:
	//
	// 2
	CreateMemberCount *string `json:"CreateMemberCount,omitempty" xml:"CreateMemberCount,omitempty"`
	// The ID of the global active database cluster.
	//
	// example:
	//
	// gad-rm-bp1npi2j8********
	GadInstanceName *string `json:"GadInstanceName,omitempty" xml:"GadInstanceName,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 5374xxxx
	TaskID *string `json:"TaskID,omitempty" xml:"TaskID,omitempty"`
}

func (s CreateGADInstanceResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s CreateGADInstanceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateGADInstanceResponseBodyResult) GetCreateMemberCount() *string {
	return s.CreateMemberCount
}

func (s *CreateGADInstanceResponseBodyResult) GetGadInstanceName() *string {
	return s.GadInstanceName
}

func (s *CreateGADInstanceResponseBodyResult) GetTaskID() *string {
	return s.TaskID
}

func (s *CreateGADInstanceResponseBodyResult) SetCreateMemberCount(v string) *CreateGADInstanceResponseBodyResult {
	s.CreateMemberCount = &v
	return s
}

func (s *CreateGADInstanceResponseBodyResult) SetGadInstanceName(v string) *CreateGADInstanceResponseBodyResult {
	s.GadInstanceName = &v
	return s
}

func (s *CreateGADInstanceResponseBodyResult) SetTaskID(v string) *CreateGADInstanceResponseBodyResult {
	s.TaskID = &v
	return s
}

func (s *CreateGADInstanceResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
