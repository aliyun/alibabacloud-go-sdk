// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseHttpsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CloseHttpsRequest
	GetClientToken() *string
}

type CloseHttpsRequest struct {
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B350****
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s CloseHttpsRequest) String() string {
	return dara.Prettify(s)
}

func (s CloseHttpsRequest) GoString() string {
	return s.String()
}

func (s *CloseHttpsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CloseHttpsRequest) SetClientToken(v string) *CloseHttpsRequest {
	s.ClientToken = &v
	return s
}

func (s *CloseHttpsRequest) Validate() error {
	return dara.Validate(s)
}
