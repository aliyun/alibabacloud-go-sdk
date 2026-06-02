// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserAuthConfigsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListUserAuthConfigsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListUserAuthConfigsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListUserAuthConfigsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListUserAuthConfigsResponseBody
	GetTotalCount() *int32
	SetUserAuthConfigs(v []*ListUserAuthConfigsResponseBodyUserAuthConfigs) *ListUserAuthConfigsResponseBody
	GetUserAuthConfigs() []*ListUserAuthConfigsResponseBodyUserAuthConfigs
}

type ListUserAuthConfigsResponseBody struct {
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAVY3rYiv9VoUJQSiCitgjgRBp055u+7M/ZFoi7I0NZHJj8bgHiGAwZWnCMJPepC+WQbLSjoLewJIqkMQqvaJO7M=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 35A48818-81F2-506C-891C-2296BE8DD667
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 0
	TotalCount      *int32                                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	UserAuthConfigs []*ListUserAuthConfigsResponseBodyUserAuthConfigs `json:"UserAuthConfigs,omitempty" xml:"UserAuthConfigs,omitempty" type:"Repeated"`
}

func (s ListUserAuthConfigsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUserAuthConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserAuthConfigsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListUserAuthConfigsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListUserAuthConfigsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUserAuthConfigsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListUserAuthConfigsResponseBody) GetUserAuthConfigs() []*ListUserAuthConfigsResponseBodyUserAuthConfigs {
	return s.UserAuthConfigs
}

func (s *ListUserAuthConfigsResponseBody) SetMaxResults(v int32) *ListUserAuthConfigsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListUserAuthConfigsResponseBody) SetNextToken(v string) *ListUserAuthConfigsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListUserAuthConfigsResponseBody) SetRequestId(v string) *ListUserAuthConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserAuthConfigsResponseBody) SetTotalCount(v int32) *ListUserAuthConfigsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListUserAuthConfigsResponseBody) SetUserAuthConfigs(v []*ListUserAuthConfigsResponseBodyUserAuthConfigs) *ListUserAuthConfigsResponseBody {
	s.UserAuthConfigs = v
	return s
}

func (s *ListUserAuthConfigsResponseBody) Validate() error {
	if s.UserAuthConfigs != nil {
		for _, item := range s.UserAuthConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListUserAuthConfigsResponseBodyUserAuthConfigs struct {
	// example:
	//
	// uac-xxxxxxxx
	AuthConfigId *string `json:"AuthConfigId,omitempty" xml:"AuthConfigId,omitempty"`
	// example:
	//
	// name
	AuthConfigName *string `json:"AuthConfigName,omitempty" xml:"AuthConfigName,omitempty"`
	// example:
	//
	// ApiKey
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// example:
	//
	// connector-xxxxxxxxx
	ConnectorId *string `json:"ConnectorId,omitempty" xml:"ConnectorId,omitempty"`
	// example:
	//
	// 1
	ConnectorVersion *string `json:"ConnectorVersion,omitempty" xml:"ConnectorVersion,omitempty"`
	// example:
	//
	// 1
	FlowCount *int32 `json:"FlowCount,omitempty" xml:"FlowCount,omitempty"`
	// example:
	//
	// 2026-04-01 14:22:33
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 2026-04-01 14:22:33
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
}

func (s ListUserAuthConfigsResponseBodyUserAuthConfigs) String() string {
	return dara.Prettify(s)
}

func (s ListUserAuthConfigsResponseBodyUserAuthConfigs) GoString() string {
	return s.String()
}

func (s *ListUserAuthConfigsResponseBodyUserAuthConfigs) GetAuthConfigId() *string {
	return s.AuthConfigId
}

func (s *ListUserAuthConfigsResponseBodyUserAuthConfigs) GetAuthConfigName() *string {
	return s.AuthConfigName
}

func (s *ListUserAuthConfigsResponseBodyUserAuthConfigs) GetAuthType() *string {
	return s.AuthType
}

func (s *ListUserAuthConfigsResponseBodyUserAuthConfigs) GetConnectorId() *string {
	return s.ConnectorId
}

func (s *ListUserAuthConfigsResponseBodyUserAuthConfigs) GetConnectorVersion() *string {
	return s.ConnectorVersion
}

func (s *ListUserAuthConfigsResponseBodyUserAuthConfigs) GetFlowCount() *int32 {
	return s.FlowCount
}

func (s *ListUserAuthConfigsResponseBodyUserAuthConfigs) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListUserAuthConfigsResponseBodyUserAuthConfigs) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListUserAuthConfigsResponseBodyUserAuthConfigs) SetAuthConfigId(v string) *ListUserAuthConfigsResponseBodyUserAuthConfigs {
	s.AuthConfigId = &v
	return s
}

func (s *ListUserAuthConfigsResponseBodyUserAuthConfigs) SetAuthConfigName(v string) *ListUserAuthConfigsResponseBodyUserAuthConfigs {
	s.AuthConfigName = &v
	return s
}

func (s *ListUserAuthConfigsResponseBodyUserAuthConfigs) SetAuthType(v string) *ListUserAuthConfigsResponseBodyUserAuthConfigs {
	s.AuthType = &v
	return s
}

func (s *ListUserAuthConfigsResponseBodyUserAuthConfigs) SetConnectorId(v string) *ListUserAuthConfigsResponseBodyUserAuthConfigs {
	s.ConnectorId = &v
	return s
}

func (s *ListUserAuthConfigsResponseBodyUserAuthConfigs) SetConnectorVersion(v string) *ListUserAuthConfigsResponseBodyUserAuthConfigs {
	s.ConnectorVersion = &v
	return s
}

func (s *ListUserAuthConfigsResponseBodyUserAuthConfigs) SetFlowCount(v int32) *ListUserAuthConfigsResponseBodyUserAuthConfigs {
	s.FlowCount = &v
	return s
}

func (s *ListUserAuthConfigsResponseBodyUserAuthConfigs) SetGmtCreate(v string) *ListUserAuthConfigsResponseBodyUserAuthConfigs {
	s.GmtCreate = &v
	return s
}

func (s *ListUserAuthConfigsResponseBodyUserAuthConfigs) SetGmtModified(v string) *ListUserAuthConfigsResponseBodyUserAuthConfigs {
	s.GmtModified = &v
	return s
}

func (s *ListUserAuthConfigsResponseBodyUserAuthConfigs) Validate() error {
	return dara.Validate(s)
}
