// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetYzdStsAKResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetYzdStsAKResponseBody
	GetCode() *string
	SetData(v *GetYzdStsAKResponseBodyData) *GetYzdStsAKResponseBody
	GetData() *GetYzdStsAKResponseBodyData
	SetMsg(v string) *GetYzdStsAKResponseBody
	GetMsg() *string
	SetRequestId(v string) *GetYzdStsAKResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetYzdStsAKResponseBody
	GetSuccess() *bool
}

type GetYzdStsAKResponseBody struct {
	Code      *string                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetYzdStsAKResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Msg       *string                      `json:"Msg,omitempty" xml:"Msg,omitempty"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetYzdStsAKResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetYzdStsAKResponseBody) GoString() string {
	return s.String()
}

func (s *GetYzdStsAKResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetYzdStsAKResponseBody) GetData() *GetYzdStsAKResponseBodyData {
	return s.Data
}

func (s *GetYzdStsAKResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *GetYzdStsAKResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetYzdStsAKResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetYzdStsAKResponseBody) SetCode(v string) *GetYzdStsAKResponseBody {
	s.Code = &v
	return s
}

func (s *GetYzdStsAKResponseBody) SetData(v *GetYzdStsAKResponseBodyData) *GetYzdStsAKResponseBody {
	s.Data = v
	return s
}

func (s *GetYzdStsAKResponseBody) SetMsg(v string) *GetYzdStsAKResponseBody {
	s.Msg = &v
	return s
}

func (s *GetYzdStsAKResponseBody) SetRequestId(v string) *GetYzdStsAKResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetYzdStsAKResponseBody) SetSuccess(v bool) *GetYzdStsAKResponseBody {
	s.Success = &v
	return s
}

func (s *GetYzdStsAKResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetYzdStsAKResponseBodyData struct {
	AppId               *string `json:"appId,omitempty" xml:"appId,omitempty"`
	Bucket              *string `json:"bucket,omitempty" xml:"bucket,omitempty"`
	Endpoint            *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	Id                  *string `json:"id,omitempty" xml:"id,omitempty"`
	InternalKnowledgeId *string `json:"internalKnowledgeId,omitempty" xml:"internalKnowledgeId,omitempty"`
	Path                *string `json:"path,omitempty" xml:"path,omitempty"`
	Secret              *string `json:"secret,omitempty" xml:"secret,omitempty"`
	Token               *string `json:"token,omitempty" xml:"token,omitempty"`
}

func (s GetYzdStsAKResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetYzdStsAKResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetYzdStsAKResponseBodyData) GetAppId() *string {
	return s.AppId
}

func (s *GetYzdStsAKResponseBodyData) GetBucket() *string {
	return s.Bucket
}

func (s *GetYzdStsAKResponseBodyData) GetEndpoint() *string {
	return s.Endpoint
}

func (s *GetYzdStsAKResponseBodyData) GetId() *string {
	return s.Id
}

func (s *GetYzdStsAKResponseBodyData) GetInternalKnowledgeId() *string {
	return s.InternalKnowledgeId
}

func (s *GetYzdStsAKResponseBodyData) GetPath() *string {
	return s.Path
}

func (s *GetYzdStsAKResponseBodyData) GetSecret() *string {
	return s.Secret
}

func (s *GetYzdStsAKResponseBodyData) GetToken() *string {
	return s.Token
}

func (s *GetYzdStsAKResponseBodyData) SetAppId(v string) *GetYzdStsAKResponseBodyData {
	s.AppId = &v
	return s
}

func (s *GetYzdStsAKResponseBodyData) SetBucket(v string) *GetYzdStsAKResponseBodyData {
	s.Bucket = &v
	return s
}

func (s *GetYzdStsAKResponseBodyData) SetEndpoint(v string) *GetYzdStsAKResponseBodyData {
	s.Endpoint = &v
	return s
}

func (s *GetYzdStsAKResponseBodyData) SetId(v string) *GetYzdStsAKResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetYzdStsAKResponseBodyData) SetInternalKnowledgeId(v string) *GetYzdStsAKResponseBodyData {
	s.InternalKnowledgeId = &v
	return s
}

func (s *GetYzdStsAKResponseBodyData) SetPath(v string) *GetYzdStsAKResponseBodyData {
	s.Path = &v
	return s
}

func (s *GetYzdStsAKResponseBodyData) SetSecret(v string) *GetYzdStsAKResponseBodyData {
	s.Secret = &v
	return s
}

func (s *GetYzdStsAKResponseBodyData) SetToken(v string) *GetYzdStsAKResponseBodyData {
	s.Token = &v
	return s
}

func (s *GetYzdStsAKResponseBodyData) Validate() error {
	return dara.Validate(s)
}
