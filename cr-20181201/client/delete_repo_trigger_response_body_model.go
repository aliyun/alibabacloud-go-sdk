// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRepoTriggerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteRepoTriggerResponseBody
	GetCode() *string
	SetIsSuccess(v bool) *DeleteRepoTriggerResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *DeleteRepoTriggerResponseBody
	GetRequestId() *string
}

type DeleteRepoTriggerResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- `true`: The request is successful.
	//
	// 	- `false`: The request fails.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 85180AE4-9A57-48F8-9EF9-68ECCE54B552
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRepoTriggerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRepoTriggerResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRepoTriggerResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteRepoTriggerResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *DeleteRepoTriggerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRepoTriggerResponseBody) SetCode(v string) *DeleteRepoTriggerResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteRepoTriggerResponseBody) SetIsSuccess(v bool) *DeleteRepoTriggerResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *DeleteRepoTriggerResponseBody) SetRequestId(v string) *DeleteRepoTriggerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRepoTriggerResponseBody) Validate() error {
	return dara.Validate(s)
}
