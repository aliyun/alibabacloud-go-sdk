// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAndRemoveFavoriteContentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *AddAndRemoveFavoriteContentResponseBody
	GetCode() *int32
	SetMessage(v string) *AddAndRemoveFavoriteContentResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddAndRemoveFavoriteContentResponseBody
	GetRequestId() *string
	SetResult(v bool) *AddAndRemoveFavoriteContentResponseBody
	GetResult() *bool
	SetSuccess(v string) *AddAndRemoveFavoriteContentResponseBody
	GetSuccess() *string
}

type AddAndRemoveFavoriteContentResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 121212121
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddAndRemoveFavoriteContentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddAndRemoveFavoriteContentResponseBody) GoString() string {
	return s.String()
}

func (s *AddAndRemoveFavoriteContentResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *AddAndRemoveFavoriteContentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddAndRemoveFavoriteContentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddAndRemoveFavoriteContentResponseBody) GetResult() *bool {
	return s.Result
}

func (s *AddAndRemoveFavoriteContentResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *AddAndRemoveFavoriteContentResponseBody) SetCode(v int32) *AddAndRemoveFavoriteContentResponseBody {
	s.Code = &v
	return s
}

func (s *AddAndRemoveFavoriteContentResponseBody) SetMessage(v string) *AddAndRemoveFavoriteContentResponseBody {
	s.Message = &v
	return s
}

func (s *AddAndRemoveFavoriteContentResponseBody) SetRequestId(v string) *AddAndRemoveFavoriteContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddAndRemoveFavoriteContentResponseBody) SetResult(v bool) *AddAndRemoveFavoriteContentResponseBody {
	s.Result = &v
	return s
}

func (s *AddAndRemoveFavoriteContentResponseBody) SetSuccess(v string) *AddAndRemoveFavoriteContentResponseBody {
	s.Success = &v
	return s
}

func (s *AddAndRemoveFavoriteContentResponseBody) Validate() error {
	return dara.Validate(s)
}
