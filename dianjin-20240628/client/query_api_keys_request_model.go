// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryApiKeysRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExternalUserId(v string) *QueryApiKeysRequest
	GetExternalUserId() *string
}

type QueryApiKeysRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1001
	ExternalUserId *string `json:"externalUserId,omitempty" xml:"externalUserId,omitempty"`
}

func (s QueryApiKeysRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryApiKeysRequest) GoString() string {
	return s.String()
}

func (s *QueryApiKeysRequest) GetExternalUserId() *string {
	return s.ExternalUserId
}

func (s *QueryApiKeysRequest) SetExternalUserId(v string) *QueryApiKeysRequest {
	s.ExternalUserId = &v
	return s
}

func (s *QueryApiKeysRequest) Validate() error {
	return dara.Validate(s)
}
