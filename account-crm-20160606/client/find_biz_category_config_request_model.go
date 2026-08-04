// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFindBizCategoryConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLocaleString(v string) *FindBizCategoryConfigRequest
	GetLocaleString() *string
	SetUserId(v int64) *FindBizCategoryConfigRequest
	GetUserId() *int64
}

type FindBizCategoryConfigRequest struct {
	LocaleString *string `json:"LocaleString,omitempty" xml:"LocaleString,omitempty"`
	// This parameter is required.
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s FindBizCategoryConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s FindBizCategoryConfigRequest) GoString() string {
	return s.String()
}

func (s *FindBizCategoryConfigRequest) GetLocaleString() *string {
	return s.LocaleString
}

func (s *FindBizCategoryConfigRequest) GetUserId() *int64 {
	return s.UserId
}

func (s *FindBizCategoryConfigRequest) SetLocaleString(v string) *FindBizCategoryConfigRequest {
	s.LocaleString = &v
	return s
}

func (s *FindBizCategoryConfigRequest) SetUserId(v int64) *FindBizCategoryConfigRequest {
	s.UserId = &v
	return s
}

func (s *FindBizCategoryConfigRequest) Validate() error {
	return dara.Validate(s)
}
