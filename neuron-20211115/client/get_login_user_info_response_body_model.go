// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLoginUserInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetResult(v *BucUser) *GetLoginUserInfoResponseBody
	GetResult() *BucUser
}

type GetLoginUserInfoResponseBody struct {
	Result *BucUser `json:"result,omitempty" xml:"result,omitempty"`
}

func (s GetLoginUserInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLoginUserInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetLoginUserInfoResponseBody) GetResult() *BucUser {
	return s.Result
}

func (s *GetLoginUserInfoResponseBody) SetResult(v *BucUser) *GetLoginUserInfoResponseBody {
	s.Result = v
	return s
}

func (s *GetLoginUserInfoResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}
