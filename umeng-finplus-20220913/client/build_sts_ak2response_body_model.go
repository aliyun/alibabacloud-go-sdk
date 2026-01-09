// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBuildStsAK2ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *BuildStsAK2ResponseBody
	GetCode() *string
	SetData(v *BuildStsAK2ResponseBodyData) *BuildStsAK2ResponseBody
	GetData() *BuildStsAK2ResponseBodyData
	SetMsg(v string) *BuildStsAK2ResponseBody
	GetMsg() *string
	SetRequestId(v string) *BuildStsAK2ResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *BuildStsAK2ResponseBody
	GetSuccess() *bool
}

type BuildStsAK2ResponseBody struct {
	Code      *string                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *BuildStsAK2ResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Msg       *string                      `json:"Msg,omitempty" xml:"Msg,omitempty"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BuildStsAK2ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BuildStsAK2ResponseBody) GoString() string {
	return s.String()
}

func (s *BuildStsAK2ResponseBody) GetCode() *string {
	return s.Code
}

func (s *BuildStsAK2ResponseBody) GetData() *BuildStsAK2ResponseBodyData {
	return s.Data
}

func (s *BuildStsAK2ResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *BuildStsAK2ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BuildStsAK2ResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *BuildStsAK2ResponseBody) SetCode(v string) *BuildStsAK2ResponseBody {
	s.Code = &v
	return s
}

func (s *BuildStsAK2ResponseBody) SetData(v *BuildStsAK2ResponseBodyData) *BuildStsAK2ResponseBody {
	s.Data = v
	return s
}

func (s *BuildStsAK2ResponseBody) SetMsg(v string) *BuildStsAK2ResponseBody {
	s.Msg = &v
	return s
}

func (s *BuildStsAK2ResponseBody) SetRequestId(v string) *BuildStsAK2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *BuildStsAK2ResponseBody) SetSuccess(v bool) *BuildStsAK2ResponseBody {
	s.Success = &v
	return s
}

func (s *BuildStsAK2ResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BuildStsAK2ResponseBodyData struct {
	Bucket   *string `json:"bucket,omitempty" xml:"bucket,omitempty"`
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	Id       *string `json:"id,omitempty" xml:"id,omitempty"`
	Path     *string `json:"path,omitempty" xml:"path,omitempty"`
	Secret   *string `json:"secret,omitempty" xml:"secret,omitempty"`
	Token    *string `json:"token,omitempty" xml:"token,omitempty"`
}

func (s BuildStsAK2ResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s BuildStsAK2ResponseBodyData) GoString() string {
	return s.String()
}

func (s *BuildStsAK2ResponseBodyData) GetBucket() *string {
	return s.Bucket
}

func (s *BuildStsAK2ResponseBodyData) GetEndpoint() *string {
	return s.Endpoint
}

func (s *BuildStsAK2ResponseBodyData) GetId() *string {
	return s.Id
}

func (s *BuildStsAK2ResponseBodyData) GetPath() *string {
	return s.Path
}

func (s *BuildStsAK2ResponseBodyData) GetSecret() *string {
	return s.Secret
}

func (s *BuildStsAK2ResponseBodyData) GetToken() *string {
	return s.Token
}

func (s *BuildStsAK2ResponseBodyData) SetBucket(v string) *BuildStsAK2ResponseBodyData {
	s.Bucket = &v
	return s
}

func (s *BuildStsAK2ResponseBodyData) SetEndpoint(v string) *BuildStsAK2ResponseBodyData {
	s.Endpoint = &v
	return s
}

func (s *BuildStsAK2ResponseBodyData) SetId(v string) *BuildStsAK2ResponseBodyData {
	s.Id = &v
	return s
}

func (s *BuildStsAK2ResponseBodyData) SetPath(v string) *BuildStsAK2ResponseBodyData {
	s.Path = &v
	return s
}

func (s *BuildStsAK2ResponseBodyData) SetSecret(v string) *BuildStsAK2ResponseBodyData {
	s.Secret = &v
	return s
}

func (s *BuildStsAK2ResponseBodyData) SetToken(v string) *BuildStsAK2ResponseBodyData {
	s.Token = &v
	return s
}

func (s *BuildStsAK2ResponseBodyData) Validate() error {
	return dara.Validate(s)
}
