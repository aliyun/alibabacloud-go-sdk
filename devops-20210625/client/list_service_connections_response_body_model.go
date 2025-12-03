// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceConnectionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListServiceConnectionsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListServiceConnectionsResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListServiceConnectionsResponseBody
	GetRequestId() *string
	SetServiceConnections(v []*ListServiceConnectionsResponseBodyServiceConnections) *ListServiceConnectionsResponseBody
	GetServiceConnections() []*ListServiceConnectionsResponseBodyServiceConnections
	SetSuccess(v bool) *ListServiceConnectionsResponseBody
	GetSuccess() *bool
}

type ListServiceConnectionsResponseBody struct {
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// ASSDS-ASSASX-XSAXSA-XSAXSAXS
	RequestId          *string                                                 `json:"requestId,omitempty" xml:"requestId,omitempty"`
	ServiceConnections []*ListServiceConnectionsResponseBodyServiceConnections `json:"serviceConnections,omitempty" xml:"serviceConnections,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListServiceConnectionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListServiceConnectionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceConnectionsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListServiceConnectionsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListServiceConnectionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListServiceConnectionsResponseBody) GetServiceConnections() []*ListServiceConnectionsResponseBodyServiceConnections {
	return s.ServiceConnections
}

func (s *ListServiceConnectionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListServiceConnectionsResponseBody) SetErrorCode(v string) *ListServiceConnectionsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListServiceConnectionsResponseBody) SetErrorMessage(v string) *ListServiceConnectionsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListServiceConnectionsResponseBody) SetRequestId(v string) *ListServiceConnectionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServiceConnectionsResponseBody) SetServiceConnections(v []*ListServiceConnectionsResponseBodyServiceConnections) *ListServiceConnectionsResponseBody {
	s.ServiceConnections = v
	return s
}

func (s *ListServiceConnectionsResponseBody) SetSuccess(v bool) *ListServiceConnectionsResponseBody {
	s.Success = &v
	return s
}

func (s *ListServiceConnectionsResponseBody) Validate() error {
	if s.ServiceConnections != nil {
		for _, item := range s.ServiceConnections {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListServiceConnectionsResponseBodyServiceConnections struct {
	// example:
	//
	// 1586863220000
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 123
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 张三的oss服务连接
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 1212123212121212
	OwnerAccountId *int64 `json:"ownerAccountId,omitempty" xml:"ownerAccountId,omitempty"`
	// example:
	//
	// oss
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

func (s ListServiceConnectionsResponseBodyServiceConnections) String() string {
	return dara.Prettify(s)
}

func (s ListServiceConnectionsResponseBodyServiceConnections) GoString() string {
	return s.String()
}

func (s *ListServiceConnectionsResponseBodyServiceConnections) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListServiceConnectionsResponseBodyServiceConnections) GetId() *int64 {
	return s.Id
}

func (s *ListServiceConnectionsResponseBodyServiceConnections) GetName() *string {
	return s.Name
}

func (s *ListServiceConnectionsResponseBodyServiceConnections) GetOwnerAccountId() *int64 {
	return s.OwnerAccountId
}

func (s *ListServiceConnectionsResponseBodyServiceConnections) GetType() *string {
	return s.Type
}

func (s *ListServiceConnectionsResponseBodyServiceConnections) GetUuid() *string {
	return s.Uuid
}

func (s *ListServiceConnectionsResponseBodyServiceConnections) SetCreateTime(v int64) *ListServiceConnectionsResponseBodyServiceConnections {
	s.CreateTime = &v
	return s
}

func (s *ListServiceConnectionsResponseBodyServiceConnections) SetId(v int64) *ListServiceConnectionsResponseBodyServiceConnections {
	s.Id = &v
	return s
}

func (s *ListServiceConnectionsResponseBodyServiceConnections) SetName(v string) *ListServiceConnectionsResponseBodyServiceConnections {
	s.Name = &v
	return s
}

func (s *ListServiceConnectionsResponseBodyServiceConnections) SetOwnerAccountId(v int64) *ListServiceConnectionsResponseBodyServiceConnections {
	s.OwnerAccountId = &v
	return s
}

func (s *ListServiceConnectionsResponseBodyServiceConnections) SetType(v string) *ListServiceConnectionsResponseBodyServiceConnections {
	s.Type = &v
	return s
}

func (s *ListServiceConnectionsResponseBodyServiceConnections) SetUuid(v string) *ListServiceConnectionsResponseBodyServiceConnections {
	s.Uuid = &v
	return s
}

func (s *ListServiceConnectionsResponseBodyServiceConnections) Validate() error {
	return dara.Validate(s)
}
