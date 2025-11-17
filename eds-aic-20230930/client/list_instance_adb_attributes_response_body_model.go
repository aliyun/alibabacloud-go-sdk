// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceAdbAttributesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListInstanceAdbAttributesResponseBodyData) *ListInstanceAdbAttributesResponseBody
	GetData() []*ListInstanceAdbAttributesResponseBodyData
	SetMaxResults(v int32) *ListInstanceAdbAttributesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListInstanceAdbAttributesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListInstanceAdbAttributesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListInstanceAdbAttributesResponseBody
	GetTotalCount() *int64
}

type ListInstanceAdbAttributesResponseBody struct {
	Data []*ListInstanceAdbAttributesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6kU+SQXzm0H9mu/FiSc****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// CB95E410-FD1D-53C5-9F7D-93CC44D7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListInstanceAdbAttributesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceAdbAttributesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceAdbAttributesResponseBody) GetData() []*ListInstanceAdbAttributesResponseBodyData {
	return s.Data
}

func (s *ListInstanceAdbAttributesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListInstanceAdbAttributesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListInstanceAdbAttributesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstanceAdbAttributesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListInstanceAdbAttributesResponseBody) SetData(v []*ListInstanceAdbAttributesResponseBodyData) *ListInstanceAdbAttributesResponseBody {
	s.Data = v
	return s
}

func (s *ListInstanceAdbAttributesResponseBody) SetMaxResults(v int32) *ListInstanceAdbAttributesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListInstanceAdbAttributesResponseBody) SetNextToken(v string) *ListInstanceAdbAttributesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListInstanceAdbAttributesResponseBody) SetRequestId(v string) *ListInstanceAdbAttributesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstanceAdbAttributesResponseBody) SetTotalCount(v int64) *ListInstanceAdbAttributesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListInstanceAdbAttributesResponseBody) Validate() error {
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

type ListInstanceAdbAttributesResponseBodyData struct {
	// example:
	//
	// 183.201.219.157
	ExternalIp *string `json:"ExternalIp,omitempty" xml:"ExternalIp,omitempty"`
	// example:
	//
	// 14840/14849
	ExternalPort *string `json:"ExternalPort,omitempty" xml:"ExternalPort,omitempty"`
	// example:
	//
	// 2024-05-15 17:33:59
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// example:
	//
	// 2024-05-15 17:33:59
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// acp-uto81vfd8t8z****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 10.0.0.239
	InternalIp *string `json:"InternalIp,omitempty" xml:"InternalIp,omitempty"`
	// example:
	//
	// 5555/5555
	InternalPort *string `json:"InternalPort,omitempty" xml:"InternalPort,omitempty"`
	// example:
	//
	// TCP
	IpProtocol *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
}

func (s ListInstanceAdbAttributesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceAdbAttributesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListInstanceAdbAttributesResponseBodyData) GetExternalIp() *string {
	return s.ExternalIp
}

func (s *ListInstanceAdbAttributesResponseBodyData) GetExternalPort() *string {
	return s.ExternalPort
}

func (s *ListInstanceAdbAttributesResponseBodyData) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *ListInstanceAdbAttributesResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListInstanceAdbAttributesResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInstanceAdbAttributesResponseBodyData) GetInternalIp() *string {
	return s.InternalIp
}

func (s *ListInstanceAdbAttributesResponseBodyData) GetInternalPort() *string {
	return s.InternalPort
}

func (s *ListInstanceAdbAttributesResponseBodyData) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *ListInstanceAdbAttributesResponseBodyData) SetExternalIp(v string) *ListInstanceAdbAttributesResponseBodyData {
	s.ExternalIp = &v
	return s
}

func (s *ListInstanceAdbAttributesResponseBodyData) SetExternalPort(v string) *ListInstanceAdbAttributesResponseBodyData {
	s.ExternalPort = &v
	return s
}

func (s *ListInstanceAdbAttributesResponseBodyData) SetGmtCreated(v string) *ListInstanceAdbAttributesResponseBodyData {
	s.GmtCreated = &v
	return s
}

func (s *ListInstanceAdbAttributesResponseBodyData) SetGmtModified(v string) *ListInstanceAdbAttributesResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *ListInstanceAdbAttributesResponseBodyData) SetInstanceId(v string) *ListInstanceAdbAttributesResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ListInstanceAdbAttributesResponseBodyData) SetInternalIp(v string) *ListInstanceAdbAttributesResponseBodyData {
	s.InternalIp = &v
	return s
}

func (s *ListInstanceAdbAttributesResponseBodyData) SetInternalPort(v string) *ListInstanceAdbAttributesResponseBodyData {
	s.InternalPort = &v
	return s
}

func (s *ListInstanceAdbAttributesResponseBodyData) SetIpProtocol(v string) *ListInstanceAdbAttributesResponseBodyData {
	s.IpProtocol = &v
	return s
}

func (s *ListInstanceAdbAttributesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
