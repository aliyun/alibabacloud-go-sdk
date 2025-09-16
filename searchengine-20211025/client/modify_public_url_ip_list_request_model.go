// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPublicUrlIpListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v map[string]*string) *ModifyPublicUrlIpListRequest
	GetBody() map[string]*string
}

type ModifyPublicUrlIpListRequest struct {
	// The request body.
	Body map[string]*string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyPublicUrlIpListRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyPublicUrlIpListRequest) GoString() string {
	return s.String()
}

func (s *ModifyPublicUrlIpListRequest) GetBody() map[string]*string {
	return s.Body
}

func (s *ModifyPublicUrlIpListRequest) SetBody(v map[string]*string) *ModifyPublicUrlIpListRequest {
	s.Body = v
	return s
}

func (s *ModifyPublicUrlIpListRequest) Validate() error {
	return dara.Validate(s)
}
