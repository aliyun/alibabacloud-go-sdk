// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceOssMountRamAuthorizeUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetInstanceOssMountRamAuthorizeUrlRequest
	GetInstanceId() *string
}

type GetInstanceOssMountRamAuthorizeUrlRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// at-cn-plt4uk4bm15
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetInstanceOssMountRamAuthorizeUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceOssMountRamAuthorizeUrlRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceOssMountRamAuthorizeUrlRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInstanceOssMountRamAuthorizeUrlRequest) SetInstanceId(v string) *GetInstanceOssMountRamAuthorizeUrlRequest {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceOssMountRamAuthorizeUrlRequest) Validate() error {
	return dara.Validate(s)
}
