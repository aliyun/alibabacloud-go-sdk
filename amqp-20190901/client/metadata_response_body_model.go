// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMetadataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *MetadataResponseBody
	GetCode() *int32
	SetData(v *MetadataResponseBodyData) *MetadataResponseBody
	GetData() *MetadataResponseBodyData
	SetMessage(v string) *MetadataResponseBody
	GetMessage() *string
	SetRequestId(v string) *MetadataResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *MetadataResponseBody
	GetSuccess() *bool
}

type MetadataResponseBody struct {
	Code      *int32                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *MetadataResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s MetadataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MetadataResponseBody) GoString() string {
	return s.String()
}

func (s *MetadataResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *MetadataResponseBody) GetData() *MetadataResponseBodyData {
	return s.Data
}

func (s *MetadataResponseBody) GetMessage() *string {
	return s.Message
}

func (s *MetadataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MetadataResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *MetadataResponseBody) SetCode(v int32) *MetadataResponseBody {
	s.Code = &v
	return s
}

func (s *MetadataResponseBody) SetData(v *MetadataResponseBodyData) *MetadataResponseBody {
	s.Data = v
	return s
}

func (s *MetadataResponseBody) SetMessage(v string) *MetadataResponseBody {
	s.Message = &v
	return s
}

func (s *MetadataResponseBody) SetRequestId(v string) *MetadataResponseBody {
	s.RequestId = &v
	return s
}

func (s *MetadataResponseBody) SetSuccess(v bool) *MetadataResponseBody {
	s.Success = &v
	return s
}

func (s *MetadataResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MetadataResponseBodyData struct {
	Endpoint             *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	HasPretendPermission *bool   `json:"HasPretendPermission,omitempty" xml:"HasPretendPermission,omitempty"`
	InternalEndpoint     *string `json:"InternalEndpoint,omitempty" xml:"InternalEndpoint,omitempty"`
	PretendUserId        *string `json:"PretendUserId,omitempty" xml:"PretendUserId,omitempty"`
	UserStatus           *int32  `json:"UserStatus,omitempty" xml:"UserStatus,omitempty"`
}

func (s MetadataResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s MetadataResponseBodyData) GoString() string {
	return s.String()
}

func (s *MetadataResponseBodyData) GetEndpoint() *string {
	return s.Endpoint
}

func (s *MetadataResponseBodyData) GetHasPretendPermission() *bool {
	return s.HasPretendPermission
}

func (s *MetadataResponseBodyData) GetInternalEndpoint() *string {
	return s.InternalEndpoint
}

func (s *MetadataResponseBodyData) GetPretendUserId() *string {
	return s.PretendUserId
}

func (s *MetadataResponseBodyData) GetUserStatus() *int32 {
	return s.UserStatus
}

func (s *MetadataResponseBodyData) SetEndpoint(v string) *MetadataResponseBodyData {
	s.Endpoint = &v
	return s
}

func (s *MetadataResponseBodyData) SetHasPretendPermission(v bool) *MetadataResponseBodyData {
	s.HasPretendPermission = &v
	return s
}

func (s *MetadataResponseBodyData) SetInternalEndpoint(v string) *MetadataResponseBodyData {
	s.InternalEndpoint = &v
	return s
}

func (s *MetadataResponseBodyData) SetPretendUserId(v string) *MetadataResponseBodyData {
	s.PretendUserId = &v
	return s
}

func (s *MetadataResponseBodyData) SetUserStatus(v int32) *MetadataResponseBodyData {
	s.UserStatus = &v
	return s
}

func (s *MetadataResponseBodyData) Validate() error {
	return dara.Validate(s)
}
