// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMessageAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetMessageAppResponseBody
	GetRequestId() *string
	SetResult(v *GetMessageAppResponseBodyResult) *GetMessageAppResponseBody
	GetResult() *GetMessageAppResponseBodyResult
}

type GetMessageAppResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 16A96B9A-****-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result.
	Result *GetMessageAppResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetMessageAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMessageAppResponseBody) GoString() string {
	return s.String()
}

func (s *GetMessageAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMessageAppResponseBody) GetResult() *GetMessageAppResponseBodyResult {
	return s.Result
}

func (s *GetMessageAppResponseBody) SetRequestId(v string) *GetMessageAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMessageAppResponseBody) SetResult(v *GetMessageAppResponseBodyResult) *GetMessageAppResponseBody {
	s.Result = v
	return s
}

func (s *GetMessageAppResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMessageAppResponseBodyResult struct {
	// The configurations of the application.
	AppConfig map[string]*string `json:"AppConfig,omitempty" xml:"AppConfig,omitempty"`
	// The ID of the interactive messaging application.
	//
	// example:
	//
	// VKL3***
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the interactive messaging application.
	//
	// example:
	//
	// testApp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The time when the interactive messaging application was created. The time is displayed in UTC.
	//
	// example:
	//
	// 502280113
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The extended field.
	Extension map[string]*string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// The status of the interactive message application. A value of 1 indicates that the application is normal.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetMessageAppResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetMessageAppResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetMessageAppResponseBodyResult) GetAppConfig() map[string]*string {
	return s.AppConfig
}

func (s *GetMessageAppResponseBodyResult) GetAppId() *string {
	return s.AppId
}

func (s *GetMessageAppResponseBodyResult) GetAppName() *string {
	return s.AppName
}

func (s *GetMessageAppResponseBodyResult) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetMessageAppResponseBodyResult) GetExtension() map[string]*string {
	return s.Extension
}

func (s *GetMessageAppResponseBodyResult) GetStatus() *int32 {
	return s.Status
}

func (s *GetMessageAppResponseBodyResult) SetAppConfig(v map[string]*string) *GetMessageAppResponseBodyResult {
	s.AppConfig = v
	return s
}

func (s *GetMessageAppResponseBodyResult) SetAppId(v string) *GetMessageAppResponseBodyResult {
	s.AppId = &v
	return s
}

func (s *GetMessageAppResponseBodyResult) SetAppName(v string) *GetMessageAppResponseBodyResult {
	s.AppName = &v
	return s
}

func (s *GetMessageAppResponseBodyResult) SetCreateTime(v int64) *GetMessageAppResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *GetMessageAppResponseBodyResult) SetExtension(v map[string]*string) *GetMessageAppResponseBodyResult {
	s.Extension = v
	return s
}

func (s *GetMessageAppResponseBodyResult) SetStatus(v int32) *GetMessageAppResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetMessageAppResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
