// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRetcodeLogstoreResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetRetcodeLogstoreResponseBodyData) *GetRetcodeLogstoreResponseBody
	GetData() *GetRetcodeLogstoreResponseBodyData
	SetRequestId(v string) *GetRetcodeLogstoreResponseBody
	GetRequestId() *string
}

type GetRetcodeLogstoreResponseBody struct {
	// The returned struct.
	Data *GetRetcodeLogstoreResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 1A9C645C-C83F-4C9D-8CCB-29BEC9E1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetRetcodeLogstoreResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRetcodeLogstoreResponseBody) GoString() string {
	return s.String()
}

func (s *GetRetcodeLogstoreResponseBody) GetData() *GetRetcodeLogstoreResponseBodyData {
	return s.Data
}

func (s *GetRetcodeLogstoreResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRetcodeLogstoreResponseBody) SetData(v *GetRetcodeLogstoreResponseBodyData) *GetRetcodeLogstoreResponseBody {
	s.Data = v
	return s
}

func (s *GetRetcodeLogstoreResponseBody) SetRequestId(v string) *GetRetcodeLogstoreResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRetcodeLogstoreResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetRetcodeLogstoreResponseBodyData struct {
	// The content of the log.
	//
	// example:
	//
	// retcode app or task can not be found!
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The information about Log Service.
	RetcodeSLSConfig *GetRetcodeLogstoreResponseBodyDataRetcodeSLSConfig `json:"RetcodeSLSConfig,omitempty" xml:"RetcodeSLSConfig,omitempty" type:"Struct"`
	// The status of the request.
	//
	// example:
	//
	// true
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetRetcodeLogstoreResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetRetcodeLogstoreResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRetcodeLogstoreResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *GetRetcodeLogstoreResponseBodyData) GetRetcodeSLSConfig() *GetRetcodeLogstoreResponseBodyDataRetcodeSLSConfig {
	return s.RetcodeSLSConfig
}

func (s *GetRetcodeLogstoreResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetRetcodeLogstoreResponseBodyData) SetMessage(v string) *GetRetcodeLogstoreResponseBodyData {
	s.Message = &v
	return s
}

func (s *GetRetcodeLogstoreResponseBodyData) SetRetcodeSLSConfig(v *GetRetcodeLogstoreResponseBodyDataRetcodeSLSConfig) *GetRetcodeLogstoreResponseBodyData {
	s.RetcodeSLSConfig = v
	return s
}

func (s *GetRetcodeLogstoreResponseBodyData) SetStatus(v string) *GetRetcodeLogstoreResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetRetcodeLogstoreResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetRetcodeLogstoreResponseBodyDataRetcodeSLSConfig struct {
	// The Log Service Logstore.
	//
	// example:
	//
	// log-test-220431
	Logstore *string `json:"Logstore,omitempty" xml:"Logstore,omitempty"`
	// The Log Service project.
	//
	// example:
	//
	// test-project
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetRetcodeLogstoreResponseBodyDataRetcodeSLSConfig) String() string {
	return dara.Prettify(s)
}

func (s GetRetcodeLogstoreResponseBodyDataRetcodeSLSConfig) GoString() string {
	return s.String()
}

func (s *GetRetcodeLogstoreResponseBodyDataRetcodeSLSConfig) GetLogstore() *string {
	return s.Logstore
}

func (s *GetRetcodeLogstoreResponseBodyDataRetcodeSLSConfig) GetProject() *string {
	return s.Project
}

func (s *GetRetcodeLogstoreResponseBodyDataRetcodeSLSConfig) GetRegionId() *string {
	return s.RegionId
}

func (s *GetRetcodeLogstoreResponseBodyDataRetcodeSLSConfig) SetLogstore(v string) *GetRetcodeLogstoreResponseBodyDataRetcodeSLSConfig {
	s.Logstore = &v
	return s
}

func (s *GetRetcodeLogstoreResponseBodyDataRetcodeSLSConfig) SetProject(v string) *GetRetcodeLogstoreResponseBodyDataRetcodeSLSConfig {
	s.Project = &v
	return s
}

func (s *GetRetcodeLogstoreResponseBodyDataRetcodeSLSConfig) SetRegionId(v string) *GetRetcodeLogstoreResponseBodyDataRetcodeSLSConfig {
	s.RegionId = &v
	return s
}

func (s *GetRetcodeLogstoreResponseBodyDataRetcodeSLSConfig) Validate() error {
	return dara.Validate(s)
}
