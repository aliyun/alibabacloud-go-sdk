// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIntegrationVersionForCSResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIntegrationVersion(v string) *GetIntegrationVersionForCSResponseBody
	GetIntegrationVersion() *string
	SetRequestId(v string) *GetIntegrationVersionForCSResponseBody
	GetRequestId() *string
}

type GetIntegrationVersionForCSResponseBody struct {
	// example:
	//
	// V1
	IntegrationVersion *string `json:"integrationVersion,omitempty" xml:"integrationVersion,omitempty"`
	// Id of the request
	//
	// example:
	//
	// CD9BCF34-EA09-5643-BC11-AF41C8DFAE5A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetIntegrationVersionForCSResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetIntegrationVersionForCSResponseBody) GoString() string {
	return s.String()
}

func (s *GetIntegrationVersionForCSResponseBody) GetIntegrationVersion() *string {
	return s.IntegrationVersion
}

func (s *GetIntegrationVersionForCSResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetIntegrationVersionForCSResponseBody) SetIntegrationVersion(v string) *GetIntegrationVersionForCSResponseBody {
	s.IntegrationVersion = &v
	return s
}

func (s *GetIntegrationVersionForCSResponseBody) SetRequestId(v string) *GetIntegrationVersionForCSResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetIntegrationVersionForCSResponseBody) Validate() error {
	return dara.Validate(s)
}
