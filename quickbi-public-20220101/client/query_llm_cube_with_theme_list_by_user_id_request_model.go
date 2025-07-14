// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryLlmCubeWithThemeListByUserIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserId(v string) *QueryLlmCubeWithThemeListByUserIdRequest
	GetUserId() *string
}

type QueryLlmCubeWithThemeListByUserIdRequest struct {
	// User ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// adsdasd-***********-123wdasd
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QueryLlmCubeWithThemeListByUserIdRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryLlmCubeWithThemeListByUserIdRequest) GoString() string {
	return s.String()
}

func (s *QueryLlmCubeWithThemeListByUserIdRequest) GetUserId() *string {
	return s.UserId
}

func (s *QueryLlmCubeWithThemeListByUserIdRequest) SetUserId(v string) *QueryLlmCubeWithThemeListByUserIdRequest {
	s.UserId = &v
	return s
}

func (s *QueryLlmCubeWithThemeListByUserIdRequest) Validate() error {
	return dara.Validate(s)
}
