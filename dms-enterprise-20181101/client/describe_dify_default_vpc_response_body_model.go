// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDifyDefaultVpcResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeDifyDefaultVpcResponseBody
	GetCode() *string
	SetData(v *DescribeDifyDefaultVpcResponseBodyData) *DescribeDifyDefaultVpcResponseBody
	GetData() *DescribeDifyDefaultVpcResponseBodyData
	SetErrorCode(v string) *DescribeDifyDefaultVpcResponseBody
	GetErrorCode() *string
	SetHttpStatusCode(v int32) *DescribeDifyDefaultVpcResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DescribeDifyDefaultVpcResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeDifyDefaultVpcResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeDifyDefaultVpcResponseBody
	GetSuccess() *bool
}

type DescribeDifyDefaultVpcResponseBody struct {
	Code           *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *DescribeDifyDefaultVpcResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode      *string                                 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	HttpStatusCode *int32                                  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDifyDefaultVpcResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDifyDefaultVpcResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDifyDefaultVpcResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeDifyDefaultVpcResponseBody) GetData() *DescribeDifyDefaultVpcResponseBodyData {
	return s.Data
}

func (s *DescribeDifyDefaultVpcResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeDifyDefaultVpcResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeDifyDefaultVpcResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeDifyDefaultVpcResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDifyDefaultVpcResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeDifyDefaultVpcResponseBody) SetCode(v string) *DescribeDifyDefaultVpcResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeDifyDefaultVpcResponseBody) SetData(v *DescribeDifyDefaultVpcResponseBodyData) *DescribeDifyDefaultVpcResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDifyDefaultVpcResponseBody) SetErrorCode(v string) *DescribeDifyDefaultVpcResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeDifyDefaultVpcResponseBody) SetHttpStatusCode(v int32) *DescribeDifyDefaultVpcResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeDifyDefaultVpcResponseBody) SetMessage(v string) *DescribeDifyDefaultVpcResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeDifyDefaultVpcResponseBody) SetRequestId(v string) *DescribeDifyDefaultVpcResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDifyDefaultVpcResponseBody) SetSuccess(v bool) *DescribeDifyDefaultVpcResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDifyDefaultVpcResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDifyDefaultVpcResponseBodyData struct {
	DefaultVpcId *string `json:"DefaultVpcId,omitempty" xml:"DefaultVpcId,omitempty"`
	WorkspaceId  *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DescribeDifyDefaultVpcResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDifyDefaultVpcResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDifyDefaultVpcResponseBodyData) GetDefaultVpcId() *string {
	return s.DefaultVpcId
}

func (s *DescribeDifyDefaultVpcResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DescribeDifyDefaultVpcResponseBodyData) SetDefaultVpcId(v string) *DescribeDifyDefaultVpcResponseBodyData {
	s.DefaultVpcId = &v
	return s
}

func (s *DescribeDifyDefaultVpcResponseBodyData) SetWorkspaceId(v string) *DescribeDifyDefaultVpcResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *DescribeDifyDefaultVpcResponseBodyData) Validate() error {
	return dara.Validate(s)
}
