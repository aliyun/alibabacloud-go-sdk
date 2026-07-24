// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSupportedConnectorsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *ListSupportedConnectorsResponseBody
	GetCode() *int64
	SetData(v []*ListSupportedConnectorsResponseBodyData) *ListSupportedConnectorsResponseBody
	GetData() []*ListSupportedConnectorsResponseBodyData
	SetMaxResults(v int32) *ListSupportedConnectorsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListSupportedConnectorsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListSupportedConnectorsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListSupportedConnectorsResponseBody
	GetSuccess() *bool
}

type ListSupportedConnectorsResponseBody struct {
	Code       *int64                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data       []*ListSupportedConnectorsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	MaxResults *int32                                     `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string                                    `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId  *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListSupportedConnectorsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSupportedConnectorsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSupportedConnectorsResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *ListSupportedConnectorsResponseBody) GetData() []*ListSupportedConnectorsResponseBodyData {
	return s.Data
}

func (s *ListSupportedConnectorsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListSupportedConnectorsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSupportedConnectorsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSupportedConnectorsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListSupportedConnectorsResponseBody) SetCode(v int64) *ListSupportedConnectorsResponseBody {
	s.Code = &v
	return s
}

func (s *ListSupportedConnectorsResponseBody) SetData(v []*ListSupportedConnectorsResponseBodyData) *ListSupportedConnectorsResponseBody {
	s.Data = v
	return s
}

func (s *ListSupportedConnectorsResponseBody) SetMaxResults(v int32) *ListSupportedConnectorsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListSupportedConnectorsResponseBody) SetNextToken(v string) *ListSupportedConnectorsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListSupportedConnectorsResponseBody) SetRequestId(v string) *ListSupportedConnectorsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSupportedConnectorsResponseBody) SetSuccess(v bool) *ListSupportedConnectorsResponseBody {
	s.Success = &v
	return s
}

func (s *ListSupportedConnectorsResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSupportedConnectorsResponseBodyData struct {
	Connector   *string `json:"Connector,omitempty" xml:"Connector,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	IconUrl     *string `json:"IconUrl,omitempty" xml:"IconUrl,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	SinkSql     *string `json:"SinkSql,omitempty" xml:"SinkSql,omitempty"`
	SourceSql   *string `json:"SourceSql,omitempty" xml:"SourceSql,omitempty"`
}

func (s ListSupportedConnectorsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListSupportedConnectorsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSupportedConnectorsResponseBodyData) GetConnector() *string {
	return s.Connector
}

func (s *ListSupportedConnectorsResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *ListSupportedConnectorsResponseBodyData) GetIconUrl() *string {
	return s.IconUrl
}

func (s *ListSupportedConnectorsResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListSupportedConnectorsResponseBodyData) GetSinkSql() *string {
	return s.SinkSql
}

func (s *ListSupportedConnectorsResponseBodyData) GetSourceSql() *string {
	return s.SourceSql
}

func (s *ListSupportedConnectorsResponseBodyData) SetConnector(v string) *ListSupportedConnectorsResponseBodyData {
	s.Connector = &v
	return s
}

func (s *ListSupportedConnectorsResponseBodyData) SetDescription(v string) *ListSupportedConnectorsResponseBodyData {
	s.Description = &v
	return s
}

func (s *ListSupportedConnectorsResponseBodyData) SetIconUrl(v string) *ListSupportedConnectorsResponseBodyData {
	s.IconUrl = &v
	return s
}

func (s *ListSupportedConnectorsResponseBodyData) SetName(v string) *ListSupportedConnectorsResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListSupportedConnectorsResponseBodyData) SetSinkSql(v string) *ListSupportedConnectorsResponseBodyData {
	s.SinkSql = &v
	return s
}

func (s *ListSupportedConnectorsResponseBodyData) SetSourceSql(v string) *ListSupportedConnectorsResponseBodyData {
	s.SourceSql = &v
	return s
}

func (s *ListSupportedConnectorsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
