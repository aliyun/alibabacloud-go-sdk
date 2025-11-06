// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteContactTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegistrantProfileIds(v string) *DeleteContactTemplatesRequest
	GetRegistrantProfileIds() *string
	SetUserClientIp(v string) *DeleteContactTemplatesRequest
	GetUserClientIp() *string
}

type DeleteContactTemplatesRequest struct {
	// This parameter is required.
	RegistrantProfileIds *string `json:"RegistrantProfileIds,omitempty" xml:"RegistrantProfileIds,omitempty"`
	// example:
	//
	// 127.0.0.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s DeleteContactTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteContactTemplatesRequest) GoString() string {
	return s.String()
}

func (s *DeleteContactTemplatesRequest) GetRegistrantProfileIds() *string {
	return s.RegistrantProfileIds
}

func (s *DeleteContactTemplatesRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *DeleteContactTemplatesRequest) SetRegistrantProfileIds(v string) *DeleteContactTemplatesRequest {
	s.RegistrantProfileIds = &v
	return s
}

func (s *DeleteContactTemplatesRequest) SetUserClientIp(v string) *DeleteContactTemplatesRequest {
	s.UserClientIp = &v
	return s
}

func (s *DeleteContactTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
