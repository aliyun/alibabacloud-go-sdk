// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTensorboardSharedUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExpireTimeSeconds(v string) *GetTensorboardSharedUrlRequest
	GetExpireTimeSeconds() *string
}

type GetTensorboardSharedUrlRequest struct {
	// The validity period of the shareable link. Unit: seconds. Maximum value: 604800.
	//
	// example:
	//
	// 86400
	ExpireTimeSeconds *string `json:"ExpireTimeSeconds,omitempty" xml:"ExpireTimeSeconds,omitempty"`
}

func (s GetTensorboardSharedUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTensorboardSharedUrlRequest) GoString() string {
	return s.String()
}

func (s *GetTensorboardSharedUrlRequest) GetExpireTimeSeconds() *string {
	return s.ExpireTimeSeconds
}

func (s *GetTensorboardSharedUrlRequest) SetExpireTimeSeconds(v string) *GetTensorboardSharedUrlRequest {
	s.ExpireTimeSeconds = &v
	return s
}

func (s *GetTensorboardSharedUrlRequest) Validate() error {
	return dara.Validate(s)
}
