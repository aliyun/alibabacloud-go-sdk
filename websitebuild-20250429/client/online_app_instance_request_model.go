// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnlineAppInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *OnlineAppInstanceRequest
	GetBizId() *string
}

type OnlineAppInstanceRequest struct {
	// The business ID of the application instance.
	//
	// example:
	//
	// WS20250801154628000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
}

func (s OnlineAppInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s OnlineAppInstanceRequest) GoString() string {
	return s.String()
}

func (s *OnlineAppInstanceRequest) GetBizId() *string {
	return s.BizId
}

func (s *OnlineAppInstanceRequest) SetBizId(v string) *OnlineAppInstanceRequest {
	s.BizId = &v
	return s
}

func (s *OnlineAppInstanceRequest) Validate() error {
	return dara.Validate(s)
}
