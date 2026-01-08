// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateAppResponseBody
	GetRequestId() *string
	SetResult(v *CreateAppResponseBodyResult) *CreateAppResponseBody
	GetResult() *CreateAppResponseBodyResult
}

type CreateAppResponseBody struct {
	// example:
	//
	// 2C5DAA30-****-5181-9B87-9D6181016197
	RequestId *string                      `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *CreateAppResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s CreateAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAppResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAppResponseBody) GetResult() *CreateAppResponseBodyResult {
	return s.Result
}

func (s *CreateAppResponseBody) SetRequestId(v string) *CreateAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAppResponseBody) SetResult(v *CreateAppResponseBodyResult) *CreateAppResponseBody {
	s.Result = v
	return s
}

func (s *CreateAppResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAppResponseBodyResult struct {
	AppId *string `json:"appId,omitempty" xml:"appId,omitempty"`
	// example:
	//
	// es-serverless-cn-xxx
	InstaneId *string `json:"instaneId,omitempty" xml:"instaneId,omitempty"`
}

func (s CreateAppResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s CreateAppResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateAppResponseBodyResult) GetAppId() *string {
	return s.AppId
}

func (s *CreateAppResponseBodyResult) GetInstaneId() *string {
	return s.InstaneId
}

func (s *CreateAppResponseBodyResult) SetAppId(v string) *CreateAppResponseBodyResult {
	s.AppId = &v
	return s
}

func (s *CreateAppResponseBodyResult) SetInstaneId(v string) *CreateAppResponseBodyResult {
	s.InstaneId = &v
	return s
}

func (s *CreateAppResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
