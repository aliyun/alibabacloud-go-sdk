// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccessKeyLastUsedInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessKey(v string) *GetAccessKeyLastUsedInfoRequest
	GetAccessKey() *string
}

type GetAccessKeyLastUsedInfoRequest struct {
	// The AccessKey ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// LTAI****************
	AccessKey *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
}

func (s GetAccessKeyLastUsedInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAccessKeyLastUsedInfoRequest) GoString() string {
	return s.String()
}

func (s *GetAccessKeyLastUsedInfoRequest) GetAccessKey() *string {
	return s.AccessKey
}

func (s *GetAccessKeyLastUsedInfoRequest) SetAccessKey(v string) *GetAccessKeyLastUsedInfoRequest {
	s.AccessKey = &v
	return s
}

func (s *GetAccessKeyLastUsedInfoRequest) Validate() error {
	return dara.Validate(s)
}
