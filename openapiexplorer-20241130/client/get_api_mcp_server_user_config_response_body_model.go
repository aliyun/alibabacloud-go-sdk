// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApiMcpServerUserConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *GetApiMcpServerUserConfigResponseBody
	GetAccountId() *string
	SetEnablePublicAccess(v bool) *GetApiMcpServerUserConfigResponseBody
	GetEnablePublicAccess() *bool
	SetGmtCreate(v string) *GetApiMcpServerUserConfigResponseBody
	GetGmtCreate() *string
	SetGmtModified(v string) *GetApiMcpServerUserConfigResponseBody
	GetGmtModified() *string
	SetRequestId(v string) *GetApiMcpServerUserConfigResponseBody
	GetRequestId() *string
	SetVpcWhitelists(v []*string) *GetApiMcpServerUserConfigResponseBody
	GetVpcWhitelists() []*string
}

type GetApiMcpServerUserConfigResponseBody struct {
	// example:
	//
	// 162302724684579*
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// example:
	//
	// true
	EnablePublicAccess *bool `json:"enablePublicAccess,omitempty" xml:"enablePublicAccess,omitempty"`
	// example:
	//
	// 2025-11-10T06:58:39Z
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 2025-11-10T06:58:39Z
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// example:
	//
	// A707AFA8-1A4C-5B2A-A165-8436C1EA38DB
	RequestId     *string   `json:"requestId,omitempty" xml:"requestId,omitempty"`
	VpcWhitelists []*string `json:"vpcWhitelists,omitempty" xml:"vpcWhitelists,omitempty" type:"Repeated"`
}

func (s GetApiMcpServerUserConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetApiMcpServerUserConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetApiMcpServerUserConfigResponseBody) GetAccountId() *string {
	return s.AccountId
}

func (s *GetApiMcpServerUserConfigResponseBody) GetEnablePublicAccess() *bool {
	return s.EnablePublicAccess
}

func (s *GetApiMcpServerUserConfigResponseBody) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetApiMcpServerUserConfigResponseBody) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetApiMcpServerUserConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetApiMcpServerUserConfigResponseBody) GetVpcWhitelists() []*string {
	return s.VpcWhitelists
}

func (s *GetApiMcpServerUserConfigResponseBody) SetAccountId(v string) *GetApiMcpServerUserConfigResponseBody {
	s.AccountId = &v
	return s
}

func (s *GetApiMcpServerUserConfigResponseBody) SetEnablePublicAccess(v bool) *GetApiMcpServerUserConfigResponseBody {
	s.EnablePublicAccess = &v
	return s
}

func (s *GetApiMcpServerUserConfigResponseBody) SetGmtCreate(v string) *GetApiMcpServerUserConfigResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *GetApiMcpServerUserConfigResponseBody) SetGmtModified(v string) *GetApiMcpServerUserConfigResponseBody {
	s.GmtModified = &v
	return s
}

func (s *GetApiMcpServerUserConfigResponseBody) SetRequestId(v string) *GetApiMcpServerUserConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetApiMcpServerUserConfigResponseBody) SetVpcWhitelists(v []*string) *GetApiMcpServerUserConfigResponseBody {
	s.VpcWhitelists = v
	return s
}

func (s *GetApiMcpServerUserConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
