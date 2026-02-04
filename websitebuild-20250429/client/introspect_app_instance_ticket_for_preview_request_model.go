// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntrospectAppInstanceTicketForPreviewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *IntrospectAppInstanceTicketForPreviewRequest
	GetBizId() *string
	SetToken(v string) *IntrospectAppInstanceTicketForPreviewRequest
	GetToken() *string
}

type IntrospectAppInstanceTicketForPreviewRequest struct {
	// example:
	//
	// WS12345678
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// example:
	//
	// ogxMqT04nnZOqdIqJZldbm-ZNsAVDz5mcqdCSudfX0SL61jjyWfV-ZnAO-OVpCt_aDl8xaaIO1OVkuvEcMn-HR_QddvaxMqIdsMY1cHdD4SDiRfOBGNdnpSdX9gG_Hi_.ab9c10816d913b51d87520ce0a72b2970668144c552689e9d9e211eb4529f0ea
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s IntrospectAppInstanceTicketForPreviewRequest) String() string {
	return dara.Prettify(s)
}

func (s IntrospectAppInstanceTicketForPreviewRequest) GoString() string {
	return s.String()
}

func (s *IntrospectAppInstanceTicketForPreviewRequest) GetBizId() *string {
	return s.BizId
}

func (s *IntrospectAppInstanceTicketForPreviewRequest) GetToken() *string {
	return s.Token
}

func (s *IntrospectAppInstanceTicketForPreviewRequest) SetBizId(v string) *IntrospectAppInstanceTicketForPreviewRequest {
	s.BizId = &v
	return s
}

func (s *IntrospectAppInstanceTicketForPreviewRequest) SetToken(v string) *IntrospectAppInstanceTicketForPreviewRequest {
	s.Token = &v
	return s
}

func (s *IntrospectAppInstanceTicketForPreviewRequest) Validate() error {
	return dara.Validate(s)
}
