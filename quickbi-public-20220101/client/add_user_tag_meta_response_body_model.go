// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddUserTagMetaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddUserTagMetaResponseBody
	GetRequestId() *string
	SetResult(v string) *AddUserTagMetaResponseBody
	GetResult() *string
	SetSuccess(v bool) *AddUserTagMetaResponseBody
	GetSuccess() *bool
}

type AddUserTagMetaResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// D787E1A3-A93C-424A-B626-C2B05DF8D885
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returns the ID of the successfully added tag.
	//
	// example:
	//
	// 0822a7d9-****-****-****-f20163ab9b0d
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// Indicates whether the request was successful. Value range:
	//
	// - true: The request succeeded
	//
	// - false: The request failed
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddUserTagMetaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddUserTagMetaResponseBody) GoString() string {
	return s.String()
}

func (s *AddUserTagMetaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddUserTagMetaResponseBody) GetResult() *string {
	return s.Result
}

func (s *AddUserTagMetaResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddUserTagMetaResponseBody) SetRequestId(v string) *AddUserTagMetaResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddUserTagMetaResponseBody) SetResult(v string) *AddUserTagMetaResponseBody {
	s.Result = &v
	return s
}

func (s *AddUserTagMetaResponseBody) SetSuccess(v bool) *AddUserTagMetaResponseBody {
	s.Success = &v
	return s
}

func (s *AddUserTagMetaResponseBody) Validate() error {
	return dara.Validate(s)
}
