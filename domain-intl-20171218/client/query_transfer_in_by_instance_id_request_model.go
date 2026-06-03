// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTransferInByInstanceIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *QueryTransferInByInstanceIdRequest
	GetInstanceId() *string
	SetLang(v string) *QueryTransferInByInstanceIdRequest
	GetLang() *string
	SetUserClientIp(v string) *QueryTransferInByInstanceIdRequest
	GetUserClientIp() *string
}

type QueryTransferInByInstanceIdRequest struct {
	// This parameter is required.
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s QueryTransferInByInstanceIdRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryTransferInByInstanceIdRequest) GoString() string {
	return s.String()
}

func (s *QueryTransferInByInstanceIdRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryTransferInByInstanceIdRequest) GetLang() *string {
	return s.Lang
}

func (s *QueryTransferInByInstanceIdRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *QueryTransferInByInstanceIdRequest) SetInstanceId(v string) *QueryTransferInByInstanceIdRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryTransferInByInstanceIdRequest) SetLang(v string) *QueryTransferInByInstanceIdRequest {
	s.Lang = &v
	return s
}

func (s *QueryTransferInByInstanceIdRequest) SetUserClientIp(v string) *QueryTransferInByInstanceIdRequest {
	s.UserClientIp = &v
	return s
}

func (s *QueryTransferInByInstanceIdRequest) Validate() error {
	return dara.Validate(s)
}
