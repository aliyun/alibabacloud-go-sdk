// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGatewayDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteGatewayDomainResponseBody
	GetCode() *int32
	SetData(v *DeleteGatewayDomainResponseBodyData) *DeleteGatewayDomainResponseBody
	GetData() *DeleteGatewayDomainResponseBodyData
	SetHttpStatusCode(v int32) *DeleteGatewayDomainResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteGatewayDomainResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteGatewayDomainResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteGatewayDomainResponseBody
	GetSuccess() *bool
}

type DeleteGatewayDomainResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *DeleteGatewayDomainResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// You are not authorized to perform this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 39016EAC-6EDB-52FE-AE20-4B013DF236FF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`: The request was successful.
	//
	// 	- `false`: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteGatewayDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteGatewayDomainResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGatewayDomainResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteGatewayDomainResponseBody) GetData() *DeleteGatewayDomainResponseBodyData {
	return s.Data
}

func (s *DeleteGatewayDomainResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteGatewayDomainResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteGatewayDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteGatewayDomainResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteGatewayDomainResponseBody) SetCode(v int32) *DeleteGatewayDomainResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteGatewayDomainResponseBody) SetData(v *DeleteGatewayDomainResponseBodyData) *DeleteGatewayDomainResponseBody {
	s.Data = v
	return s
}

func (s *DeleteGatewayDomainResponseBody) SetHttpStatusCode(v int32) *DeleteGatewayDomainResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteGatewayDomainResponseBody) SetMessage(v string) *DeleteGatewayDomainResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteGatewayDomainResponseBody) SetRequestId(v string) *DeleteGatewayDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteGatewayDomainResponseBody) SetSuccess(v bool) *DeleteGatewayDomainResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteGatewayDomainResponseBody) Validate() error {
	return dara.Validate(s)
}

type DeleteGatewayDomainResponseBodyData struct {
	// The ID of the certificate.
	//
	// example:
	//
	// 243
	CertIdentifier *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
	// The ID of the gateway.
	//
	// example:
	//
	// 253
	GatewayId *int64 `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// The unique ID of the gateway.
	//
	// example:
	//
	// gw-7ea3da97b96543e19f6c597c****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The time when the domain name was added.
	//
	// example:
	//
	// 2022-01-07T10:07:57.000+0000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The last modification time.
	//
	// example:
	//
	// 2022-01-07T10:07:57.000+0000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The domain ID.
	//
	// example:
	//
	// 2
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Indicates whether HTTP probing is allowed.
	//
	// example:
	//
	// true
	MustHttps *bool `json:"MustHttps,omitempty" xml:"MustHttps,omitempty"`
	// The name.
	//
	// example:
	//
	// CONTACTINFO
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The protocol.
	//
	// example:
	//
	// HTTP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
}

func (s DeleteGatewayDomainResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteGatewayDomainResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteGatewayDomainResponseBodyData) GetCertIdentifier() *string {
	return s.CertIdentifier
}

func (s *DeleteGatewayDomainResponseBodyData) GetGatewayId() *int64 {
	return s.GatewayId
}

func (s *DeleteGatewayDomainResponseBodyData) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *DeleteGatewayDomainResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *DeleteGatewayDomainResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DeleteGatewayDomainResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *DeleteGatewayDomainResponseBodyData) GetMustHttps() *bool {
	return s.MustHttps
}

func (s *DeleteGatewayDomainResponseBodyData) GetName() *string {
	return s.Name
}

func (s *DeleteGatewayDomainResponseBodyData) GetProtocol() *string {
	return s.Protocol
}

func (s *DeleteGatewayDomainResponseBodyData) SetCertIdentifier(v string) *DeleteGatewayDomainResponseBodyData {
	s.CertIdentifier = &v
	return s
}

func (s *DeleteGatewayDomainResponseBodyData) SetGatewayId(v int64) *DeleteGatewayDomainResponseBodyData {
	s.GatewayId = &v
	return s
}

func (s *DeleteGatewayDomainResponseBodyData) SetGatewayUniqueId(v string) *DeleteGatewayDomainResponseBodyData {
	s.GatewayUniqueId = &v
	return s
}

func (s *DeleteGatewayDomainResponseBodyData) SetGmtCreate(v string) *DeleteGatewayDomainResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *DeleteGatewayDomainResponseBodyData) SetGmtModified(v string) *DeleteGatewayDomainResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *DeleteGatewayDomainResponseBodyData) SetId(v int64) *DeleteGatewayDomainResponseBodyData {
	s.Id = &v
	return s
}

func (s *DeleteGatewayDomainResponseBodyData) SetMustHttps(v bool) *DeleteGatewayDomainResponseBodyData {
	s.MustHttps = &v
	return s
}

func (s *DeleteGatewayDomainResponseBodyData) SetName(v string) *DeleteGatewayDomainResponseBodyData {
	s.Name = &v
	return s
}

func (s *DeleteGatewayDomainResponseBodyData) SetProtocol(v string) *DeleteGatewayDomainResponseBodyData {
	s.Protocol = &v
	return s
}

func (s *DeleteGatewayDomainResponseBodyData) Validate() error {
	return dara.Validate(s)
}
