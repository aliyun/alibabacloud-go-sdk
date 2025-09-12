// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBuildStsAKResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *BuildStsAKResponseBody
	GetCode() *string
	SetData(v *BuildStsAKResponseBodyData) *BuildStsAKResponseBody
	GetData() *BuildStsAKResponseBodyData
	SetMsg(v string) *BuildStsAKResponseBody
	GetMsg() *string
	SetRequestId(v string) *BuildStsAKResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *BuildStsAKResponseBody
	GetSuccess() *bool
}

type BuildStsAKResponseBody struct {
	Code      *string                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *BuildStsAKResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Msg       *string                     `json:"Msg,omitempty" xml:"Msg,omitempty"`
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BuildStsAKResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BuildStsAKResponseBody) GoString() string {
	return s.String()
}

func (s *BuildStsAKResponseBody) GetCode() *string {
	return s.Code
}

func (s *BuildStsAKResponseBody) GetData() *BuildStsAKResponseBodyData {
	return s.Data
}

func (s *BuildStsAKResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *BuildStsAKResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BuildStsAKResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *BuildStsAKResponseBody) SetCode(v string) *BuildStsAKResponseBody {
	s.Code = &v
	return s
}

func (s *BuildStsAKResponseBody) SetData(v *BuildStsAKResponseBodyData) *BuildStsAKResponseBody {
	s.Data = v
	return s
}

func (s *BuildStsAKResponseBody) SetMsg(v string) *BuildStsAKResponseBody {
	s.Msg = &v
	return s
}

func (s *BuildStsAKResponseBody) SetRequestId(v string) *BuildStsAKResponseBody {
	s.RequestId = &v
	return s
}

func (s *BuildStsAKResponseBody) SetSuccess(v bool) *BuildStsAKResponseBody {
	s.Success = &v
	return s
}

func (s *BuildStsAKResponseBody) Validate() error {
	return dara.Validate(s)
}

type BuildStsAKResponseBodyData struct {
	Bucket   *string `json:"bucket,omitempty" xml:"bucket,omitempty"`
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	Id       *string `json:"id,omitempty" xml:"id,omitempty"`
	Path     *string `json:"path,omitempty" xml:"path,omitempty"`
	Secret   *string `json:"secret,omitempty" xml:"secret,omitempty"`
	Token    *string `json:"token,omitempty" xml:"token,omitempty"`
}

func (s BuildStsAKResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s BuildStsAKResponseBodyData) GoString() string {
	return s.String()
}

func (s *BuildStsAKResponseBodyData) GetBucket() *string {
	return s.Bucket
}

func (s *BuildStsAKResponseBodyData) GetEndpoint() *string {
	return s.Endpoint
}

func (s *BuildStsAKResponseBodyData) GetId() *string {
	return s.Id
}

func (s *BuildStsAKResponseBodyData) GetPath() *string {
	return s.Path
}

func (s *BuildStsAKResponseBodyData) GetSecret() *string {
	return s.Secret
}

func (s *BuildStsAKResponseBodyData) GetToken() *string {
	return s.Token
}

func (s *BuildStsAKResponseBodyData) SetBucket(v string) *BuildStsAKResponseBodyData {
	s.Bucket = &v
	return s
}

func (s *BuildStsAKResponseBodyData) SetEndpoint(v string) *BuildStsAKResponseBodyData {
	s.Endpoint = &v
	return s
}

func (s *BuildStsAKResponseBodyData) SetId(v string) *BuildStsAKResponseBodyData {
	s.Id = &v
	return s
}

func (s *BuildStsAKResponseBodyData) SetPath(v string) *BuildStsAKResponseBodyData {
	s.Path = &v
	return s
}

func (s *BuildStsAKResponseBodyData) SetSecret(v string) *BuildStsAKResponseBodyData {
	s.Secret = &v
	return s
}

func (s *BuildStsAKResponseBodyData) SetToken(v string) *BuildStsAKResponseBodyData {
	s.Token = &v
	return s
}

func (s *BuildStsAKResponseBodyData) Validate() error {
	return dara.Validate(s)
}
