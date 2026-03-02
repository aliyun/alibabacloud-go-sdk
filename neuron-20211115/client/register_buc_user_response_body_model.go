// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterBucUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RegisterBucUserResponseBody
	GetRequestId() *string
	SetResult(v *BucUser) *RegisterBucUserResponseBody
	GetResult() *BucUser
}

type RegisterBucUserResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// D372D265-81C4-5B84-8827-596F0CF768FF
	RequestId *string  `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *BucUser `json:"result,omitempty" xml:"result,omitempty"`
}

func (s RegisterBucUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RegisterBucUserResponseBody) GoString() string {
	return s.String()
}

func (s *RegisterBucUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RegisterBucUserResponseBody) GetResult() *BucUser {
	return s.Result
}

func (s *RegisterBucUserResponseBody) SetRequestId(v string) *RegisterBucUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *RegisterBucUserResponseBody) SetResult(v *BucUser) *RegisterBucUserResponseBody {
	s.Result = v
	return s
}

func (s *RegisterBucUserResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}
