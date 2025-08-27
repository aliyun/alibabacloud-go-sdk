// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryGroupCorpListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserId(v string) *QueryGroupCorpListRequest
	GetUserId() *string
}

type QueryGroupCorpListRequest struct {
	// example:
	//
	// 123
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s QueryGroupCorpListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryGroupCorpListRequest) GoString() string {
	return s.String()
}

func (s *QueryGroupCorpListRequest) GetUserId() *string {
	return s.UserId
}

func (s *QueryGroupCorpListRequest) SetUserId(v string) *QueryGroupCorpListRequest {
	s.UserId = &v
	return s
}

func (s *QueryGroupCorpListRequest) Validate() error {
	return dara.Validate(s)
}
