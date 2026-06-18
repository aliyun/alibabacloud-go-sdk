// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWorkspaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteWorkspaceResponseBody
	GetCode() *string
	SetFailReasons(v []*DeleteWorkspaceResponseBodyFailReasons) *DeleteWorkspaceResponseBody
	GetFailReasons() []*DeleteWorkspaceResponseBodyFailReasons
	SetHttpStatusCode(v int32) *DeleteWorkspaceResponseBody
	GetHttpStatusCode() *int32
	SetIsDeleted(v bool) *DeleteWorkspaceResponseBody
	GetIsDeleted() *bool
	SetMessage(v string) *DeleteWorkspaceResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteWorkspaceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteWorkspaceResponseBody
	GetSuccess() *bool
}

type DeleteWorkspaceResponseBody struct {
	// The response status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The list of failure reasons.
	FailReasons []*DeleteWorkspaceResponseBodyFailReasons `json:"failReasons,omitempty" xml:"failReasons,omitempty" type:"Repeated"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// Indicates whether the workspace is successfully deleted. Valid values:
	//
	// - true: Succeeded.
	//
	// - false: Failed.
	//
	// example:
	//
	// true
	IsDeleted *bool `json:"isDeleted,omitempty" xml:"isDeleted,omitempty"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// DFD55E7B-0615-5343-BDA0-FBEE86F29935
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the API call is successful. Valid values:
	//
	// - true: Succeeded.
	//
	// - false: Failed.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteWorkspaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteWorkspaceResponseBody) GetFailReasons() []*DeleteWorkspaceResponseBodyFailReasons {
	return s.FailReasons
}

func (s *DeleteWorkspaceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteWorkspaceResponseBody) GetIsDeleted() *bool {
	return s.IsDeleted
}

func (s *DeleteWorkspaceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteWorkspaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteWorkspaceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteWorkspaceResponseBody) SetCode(v string) *DeleteWorkspaceResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteWorkspaceResponseBody) SetFailReasons(v []*DeleteWorkspaceResponseBodyFailReasons) *DeleteWorkspaceResponseBody {
	s.FailReasons = v
	return s
}

func (s *DeleteWorkspaceResponseBody) SetHttpStatusCode(v int32) *DeleteWorkspaceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteWorkspaceResponseBody) SetIsDeleted(v bool) *DeleteWorkspaceResponseBody {
	s.IsDeleted = &v
	return s
}

func (s *DeleteWorkspaceResponseBody) SetMessage(v string) *DeleteWorkspaceResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteWorkspaceResponseBody) SetRequestId(v string) *DeleteWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteWorkspaceResponseBody) SetSuccess(v bool) *DeleteWorkspaceResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteWorkspaceResponseBody) Validate() error {
	if s.FailReasons != nil {
		for _, item := range s.FailReasons {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DeleteWorkspaceResponseBodyFailReasons struct {
	// The specific reason.
	//
	// example:
	//
	// API Key exists. Please clear them first.
	Reason *string `json:"reason,omitempty" xml:"reason,omitempty"`
	// The resource type.
	//
	// example:
	//
	// API Key
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s DeleteWorkspaceResponseBodyFailReasons) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkspaceResponseBodyFailReasons) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceResponseBodyFailReasons) GetReason() *string {
	return s.Reason
}

func (s *DeleteWorkspaceResponseBodyFailReasons) GetResourceType() *string {
	return s.ResourceType
}

func (s *DeleteWorkspaceResponseBodyFailReasons) SetReason(v string) *DeleteWorkspaceResponseBodyFailReasons {
	s.Reason = &v
	return s
}

func (s *DeleteWorkspaceResponseBodyFailReasons) SetResourceType(v string) *DeleteWorkspaceResponseBodyFailReasons {
	s.ResourceType = &v
	return s
}

func (s *DeleteWorkspaceResponseBodyFailReasons) Validate() error {
	return dara.Validate(s)
}
