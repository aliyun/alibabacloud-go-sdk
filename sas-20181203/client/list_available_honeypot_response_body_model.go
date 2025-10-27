// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAvailableHoneypotResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListAvailableHoneypotResponseBody
	GetCode() *string
	SetCount(v int32) *ListAvailableHoneypotResponseBody
	GetCount() *int32
	SetData(v []*ListAvailableHoneypotResponseBodyData) *ListAvailableHoneypotResponseBody
	GetData() []*ListAvailableHoneypotResponseBodyData
	SetHttpStatusCode(v int32) *ListAvailableHoneypotResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListAvailableHoneypotResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListAvailableHoneypotResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListAvailableHoneypotResponseBody
	GetSuccess() *bool
}

type ListAvailableHoneypotResponseBody struct {
	// The status code returned. The status code **200*	- indicates that the request was successful. Other status codes indicate that the request failed. You can identify the cause of the failure based on the status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The number of images that are used for the honeypot.
	//
	// example:
	//
	// 22
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// An array consisting of the information about the images that are used for the honeypot.
	Data []*ListAvailableHoneypotResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 6C24D883-984D-52FD-BB66-5F89F86E4837
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListAvailableHoneypotResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAvailableHoneypotResponseBody) GoString() string {
	return s.String()
}

func (s *ListAvailableHoneypotResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAvailableHoneypotResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *ListAvailableHoneypotResponseBody) GetData() []*ListAvailableHoneypotResponseBodyData {
	return s.Data
}

func (s *ListAvailableHoneypotResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListAvailableHoneypotResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAvailableHoneypotResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAvailableHoneypotResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAvailableHoneypotResponseBody) SetCode(v string) *ListAvailableHoneypotResponseBody {
	s.Code = &v
	return s
}

func (s *ListAvailableHoneypotResponseBody) SetCount(v int32) *ListAvailableHoneypotResponseBody {
	s.Count = &v
	return s
}

func (s *ListAvailableHoneypotResponseBody) SetData(v []*ListAvailableHoneypotResponseBodyData) *ListAvailableHoneypotResponseBody {
	s.Data = v
	return s
}

func (s *ListAvailableHoneypotResponseBody) SetHttpStatusCode(v int32) *ListAvailableHoneypotResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListAvailableHoneypotResponseBody) SetMessage(v string) *ListAvailableHoneypotResponseBody {
	s.Message = &v
	return s
}

func (s *ListAvailableHoneypotResponseBody) SetRequestId(v string) *ListAvailableHoneypotResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAvailableHoneypotResponseBody) SetSuccess(v bool) *ListAvailableHoneypotResponseBody {
	s.Success = &v
	return s
}

func (s *ListAvailableHoneypotResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAvailableHoneypotResponseBodyData struct {
	// The display name of the image.
	//
	// example:
	//
	// RuoYi
	HoneypotImageDisplayName *string `json:"HoneypotImageDisplayName,omitempty" xml:"HoneypotImageDisplayName,omitempty"`
	// The ID of the image.
	//
	// example:
	//
	// sha256:007095d6de9c7a343e9fc1f74a7efc9c5de9d5454789d2fa505a1b3fc62****
	HoneypotImageId *string `json:"HoneypotImageId,omitempty" xml:"HoneypotImageId,omitempty"`
	// The name of the image that is used for the honeypot.
	//
	// example:
	//
	// ruoyi
	HoneypotImageName *string `json:"HoneypotImageName,omitempty" xml:"HoneypotImageName,omitempty"`
	// The type of the image.
	//
	// example:
	//
	// Web
	HoneypotImageType *string `json:"HoneypotImageType,omitempty" xml:"HoneypotImageType,omitempty"`
	// The version of the image.
	//
	// example:
	//
	// 1.0.2
	HoneypotImageVersion *string `json:"HoneypotImageVersion,omitempty" xml:"HoneypotImageVersion,omitempty"`
	// The port that is supported by the honeypot. The value is in the JSON format. Valid values:
	//
	// 	- **log_type**: the log type
	//
	// 	- **proto**: the supported protocol
	//
	// 	- **description**: the description
	//
	// 	- **ports**: the supported ports
	//
	// 	- **port_str**: the supported port number of the string type
	//
	// 	- **type**: the type
	//
	// example:
	//
	// [{"log_type":"web","proto":"tcp","description":"webServerPort","ports":[80.0],"port_str":"80","type":"web"}]
	Multiports *string `json:"Multiports,omitempty" xml:"Multiports,omitempty"`
	// The protocol that is supported by the honeypot.
	//
	// example:
	//
	// tcp
	Proto *string `json:"Proto,omitempty" xml:"Proto,omitempty"`
	// The service port of the honeypot.
	//
	// example:
	//
	// 27017.0
	ServicePort *string `json:"ServicePort,omitempty" xml:"ServicePort,omitempty"`
	// The configuration template of the honeypot.
	//
	// example:
	//
	// {}
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
}

func (s ListAvailableHoneypotResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAvailableHoneypotResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAvailableHoneypotResponseBodyData) GetHoneypotImageDisplayName() *string {
	return s.HoneypotImageDisplayName
}

func (s *ListAvailableHoneypotResponseBodyData) GetHoneypotImageId() *string {
	return s.HoneypotImageId
}

func (s *ListAvailableHoneypotResponseBodyData) GetHoneypotImageName() *string {
	return s.HoneypotImageName
}

func (s *ListAvailableHoneypotResponseBodyData) GetHoneypotImageType() *string {
	return s.HoneypotImageType
}

func (s *ListAvailableHoneypotResponseBodyData) GetHoneypotImageVersion() *string {
	return s.HoneypotImageVersion
}

func (s *ListAvailableHoneypotResponseBodyData) GetMultiports() *string {
	return s.Multiports
}

func (s *ListAvailableHoneypotResponseBodyData) GetProto() *string {
	return s.Proto
}

func (s *ListAvailableHoneypotResponseBodyData) GetServicePort() *string {
	return s.ServicePort
}

func (s *ListAvailableHoneypotResponseBodyData) GetTemplate() *string {
	return s.Template
}

func (s *ListAvailableHoneypotResponseBodyData) SetHoneypotImageDisplayName(v string) *ListAvailableHoneypotResponseBodyData {
	s.HoneypotImageDisplayName = &v
	return s
}

func (s *ListAvailableHoneypotResponseBodyData) SetHoneypotImageId(v string) *ListAvailableHoneypotResponseBodyData {
	s.HoneypotImageId = &v
	return s
}

func (s *ListAvailableHoneypotResponseBodyData) SetHoneypotImageName(v string) *ListAvailableHoneypotResponseBodyData {
	s.HoneypotImageName = &v
	return s
}

func (s *ListAvailableHoneypotResponseBodyData) SetHoneypotImageType(v string) *ListAvailableHoneypotResponseBodyData {
	s.HoneypotImageType = &v
	return s
}

func (s *ListAvailableHoneypotResponseBodyData) SetHoneypotImageVersion(v string) *ListAvailableHoneypotResponseBodyData {
	s.HoneypotImageVersion = &v
	return s
}

func (s *ListAvailableHoneypotResponseBodyData) SetMultiports(v string) *ListAvailableHoneypotResponseBodyData {
	s.Multiports = &v
	return s
}

func (s *ListAvailableHoneypotResponseBodyData) SetProto(v string) *ListAvailableHoneypotResponseBodyData {
	s.Proto = &v
	return s
}

func (s *ListAvailableHoneypotResponseBodyData) SetServicePort(v string) *ListAvailableHoneypotResponseBodyData {
	s.ServicePort = &v
	return s
}

func (s *ListAvailableHoneypotResponseBodyData) SetTemplate(v string) *ListAvailableHoneypotResponseBodyData {
	s.Template = &v
	return s
}

func (s *ListAvailableHoneypotResponseBodyData) Validate() error {
	return dara.Validate(s)
}
