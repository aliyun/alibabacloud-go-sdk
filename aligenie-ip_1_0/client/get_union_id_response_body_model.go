// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUnionIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *GetUnionIdResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetUnionIdResponseBody
	GetRequestId() *string
	SetResult(v []*GetUnionIdResponseBodyResult) *GetUnionIdResponseBody
	GetResult() []*GetUnionIdResponseBodyResult
	SetStatusCode(v int32) *GetUnionIdResponseBody
	GetStatusCode() *int32
}

type GetUnionIdResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 0EC7*726E
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*GetUnionIdResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s GetUnionIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUnionIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetUnionIdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetUnionIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUnionIdResponseBody) GetResult() []*GetUnionIdResponseBodyResult {
	return s.Result
}

func (s *GetUnionIdResponseBody) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUnionIdResponseBody) SetMessage(v string) *GetUnionIdResponseBody {
	s.Message = &v
	return s
}

func (s *GetUnionIdResponseBody) SetRequestId(v string) *GetUnionIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUnionIdResponseBody) SetResult(v []*GetUnionIdResponseBodyResult) *GetUnionIdResponseBody {
	s.Result = v
	return s
}

func (s *GetUnionIdResponseBody) SetStatusCode(v int32) *GetUnionIdResponseBody {
	s.StatusCode = &v
	return s
}

func (s *GetUnionIdResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetUnionIdResponseBodyResult struct {
	// example:
	//
	// 4325***765
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	// example:
	//
	// 8bh2****8s8
	UnionId *string `json:"UnionId,omitempty" xml:"UnionId,omitempty"`
}

func (s GetUnionIdResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetUnionIdResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetUnionIdResponseBodyResult) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *GetUnionIdResponseBodyResult) GetUnionId() *string {
	return s.UnionId
}

func (s *GetUnionIdResponseBodyResult) SetOrganizationId(v string) *GetUnionIdResponseBodyResult {
	s.OrganizationId = &v
	return s
}

func (s *GetUnionIdResponseBodyResult) SetUnionId(v string) *GetUnionIdResponseBodyResult {
	s.UnionId = &v
	return s
}

func (s *GetUnionIdResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
