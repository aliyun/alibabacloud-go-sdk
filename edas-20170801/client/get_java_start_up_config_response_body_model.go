// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJavaStartUpConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetJavaStartUpConfigResponseBody
	GetCode() *int32
	SetJavaStartUpConfig(v *GetJavaStartUpConfigResponseBodyJavaStartUpConfig) *GetJavaStartUpConfigResponseBody
	GetJavaStartUpConfig() *GetJavaStartUpConfigResponseBodyJavaStartUpConfig
	SetMessage(v string) *GetJavaStartUpConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetJavaStartUpConfigResponseBody
	GetRequestId() *string
}

type GetJavaStartUpConfigResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The configuration of Java startup parameters.
	JavaStartUpConfig *GetJavaStartUpConfigResponseBodyJavaStartUpConfig `json:"JavaStartUpConfig,omitempty" xml:"JavaStartUpConfig,omitempty" type:"Struct"`
	// The message that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4823-bhjf-23u4-eiufh
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetJavaStartUpConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetJavaStartUpConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetJavaStartUpConfigResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetJavaStartUpConfigResponseBody) GetJavaStartUpConfig() *GetJavaStartUpConfigResponseBodyJavaStartUpConfig {
	return s.JavaStartUpConfig
}

func (s *GetJavaStartUpConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetJavaStartUpConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetJavaStartUpConfigResponseBody) SetCode(v int32) *GetJavaStartUpConfigResponseBody {
	s.Code = &v
	return s
}

func (s *GetJavaStartUpConfigResponseBody) SetJavaStartUpConfig(v *GetJavaStartUpConfigResponseBodyJavaStartUpConfig) *GetJavaStartUpConfigResponseBody {
	s.JavaStartUpConfig = v
	return s
}

func (s *GetJavaStartUpConfigResponseBody) SetMessage(v string) *GetJavaStartUpConfigResponseBody {
	s.Message = &v
	return s
}

func (s *GetJavaStartUpConfigResponseBody) SetRequestId(v string) *GetJavaStartUpConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetJavaStartUpConfigResponseBody) Validate() error {
	if s.JavaStartUpConfig != nil {
		if err := s.JavaStartUpConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetJavaStartUpConfigResponseBodyJavaStartUpConfig struct {
	// The displayed startup parameter configuration.
	//
	// example:
	//
	// -Xms512m
	OriginalConfigs *string `json:"OriginalConfigs,omitempty" xml:"OriginalConfigs,omitempty"`
	// The effective startup parameter configuration.
	//
	// example:
	//
	// -Xms512m
	StartUpArgs *string `json:"StartUpArgs,omitempty" xml:"StartUpArgs,omitempty"`
}

func (s GetJavaStartUpConfigResponseBodyJavaStartUpConfig) String() string {
	return dara.Prettify(s)
}

func (s GetJavaStartUpConfigResponseBodyJavaStartUpConfig) GoString() string {
	return s.String()
}

func (s *GetJavaStartUpConfigResponseBodyJavaStartUpConfig) GetOriginalConfigs() *string {
	return s.OriginalConfigs
}

func (s *GetJavaStartUpConfigResponseBodyJavaStartUpConfig) GetStartUpArgs() *string {
	return s.StartUpArgs
}

func (s *GetJavaStartUpConfigResponseBodyJavaStartUpConfig) SetOriginalConfigs(v string) *GetJavaStartUpConfigResponseBodyJavaStartUpConfig {
	s.OriginalConfigs = &v
	return s
}

func (s *GetJavaStartUpConfigResponseBodyJavaStartUpConfig) SetStartUpArgs(v string) *GetJavaStartUpConfigResponseBodyJavaStartUpConfig {
	s.StartUpArgs = &v
	return s
}

func (s *GetJavaStartUpConfigResponseBodyJavaStartUpConfig) Validate() error {
	return dara.Validate(s)
}
