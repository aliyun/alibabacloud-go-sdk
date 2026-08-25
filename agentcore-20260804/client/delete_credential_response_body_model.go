// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCredentialResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteCredentialResponseBody
	GetCode() *string
	SetData(v *DeleteCredentialResponseBodyData) *DeleteCredentialResponseBody
	GetData() *DeleteCredentialResponseBodyData
	SetHttpStatusCode(v int32) *DeleteCredentialResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteCredentialResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteCredentialResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteCredentialResponseBody
	GetSuccess() *bool
}

type DeleteCredentialResponseBody struct {
	// example:
	//
	// SUCCESS
	Code *string                           `json:"code,omitempty" xml:"code,omitempty"`
	Data *DeleteCredentialResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// request-123456
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteCredentialResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCredentialResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCredentialResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteCredentialResponseBody) GetData() *DeleteCredentialResponseBodyData {
	return s.Data
}

func (s *DeleteCredentialResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteCredentialResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteCredentialResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCredentialResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteCredentialResponseBody) SetCode(v string) *DeleteCredentialResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteCredentialResponseBody) SetData(v *DeleteCredentialResponseBodyData) *DeleteCredentialResponseBody {
	s.Data = v
	return s
}

func (s *DeleteCredentialResponseBody) SetHttpStatusCode(v int32) *DeleteCredentialResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteCredentialResponseBody) SetMessage(v string) *DeleteCredentialResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteCredentialResponseBody) SetRequestId(v string) *DeleteCredentialResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCredentialResponseBody) SetSuccess(v bool) *DeleteCredentialResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteCredentialResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteCredentialResponseBodyData struct {
	// example:
	//
	// cred-123456
	CredentialId *string `json:"credentialId,omitempty" xml:"credentialId,omitempty"`
	// example:
	//
	// model-api-key
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// ws-123456
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s DeleteCredentialResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteCredentialResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteCredentialResponseBodyData) GetCredentialId() *string {
	return s.CredentialId
}

func (s *DeleteCredentialResponseBodyData) GetName() *string {
	return s.Name
}

func (s *DeleteCredentialResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DeleteCredentialResponseBodyData) SetCredentialId(v string) *DeleteCredentialResponseBodyData {
	s.CredentialId = &v
	return s
}

func (s *DeleteCredentialResponseBodyData) SetName(v string) *DeleteCredentialResponseBodyData {
	s.Name = &v
	return s
}

func (s *DeleteCredentialResponseBodyData) SetWorkspaceId(v string) *DeleteCredentialResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteCredentialResponseBodyData) Validate() error {
	return dara.Validate(s)
}
