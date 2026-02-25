// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHTTPGetAction interface {
	dara.Model
	String() string
	GoString() string
	SetHost(v string) *HTTPGetAction
	GetHost() *string
	SetHttpHeaders(v []*HTTPHeader) *HTTPGetAction
	GetHttpHeaders() []*HTTPHeader
	SetPath(v string) *HTTPGetAction
	GetPath() *string
	SetPort(v int32) *HTTPGetAction
	GetPort() *int32
	SetScheme(v string) *HTTPGetAction
	GetScheme() *string
}

type HTTPGetAction struct {
	// The hostname to which you want to connect. The default value is the IP address of the pod. You may need to specify Host in HttpHeaders.
	//
	// example:
	//
	// 172.22.XX.XX
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The custom headers that you need to specify in the request. Duplicate headers are allowed in an HTTP request.
	HttpHeaders []*HTTPHeader `json:"HttpHeaders,omitempty" xml:"HttpHeaders,omitempty" type:"Repeated"`
	// The path of a URL.
	//
	// example:
	//
	// /path1
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The port range. Valid values: 1 to 65535.
	//
	// example:
	//
	// 80
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The scheme that you want to use to connect to the host. Default value: http.
	//
	// example:
	//
	// http
	Scheme *string `json:"Scheme,omitempty" xml:"Scheme,omitempty"`
}

func (s HTTPGetAction) String() string {
	return dara.Prettify(s)
}

func (s HTTPGetAction) GoString() string {
	return s.String()
}

func (s *HTTPGetAction) GetHost() *string {
	return s.Host
}

func (s *HTTPGetAction) GetHttpHeaders() []*HTTPHeader {
	return s.HttpHeaders
}

func (s *HTTPGetAction) GetPath() *string {
	return s.Path
}

func (s *HTTPGetAction) GetPort() *int32 {
	return s.Port
}

func (s *HTTPGetAction) GetScheme() *string {
	return s.Scheme
}

func (s *HTTPGetAction) SetHost(v string) *HTTPGetAction {
	s.Host = &v
	return s
}

func (s *HTTPGetAction) SetHttpHeaders(v []*HTTPHeader) *HTTPGetAction {
	s.HttpHeaders = v
	return s
}

func (s *HTTPGetAction) SetPath(v string) *HTTPGetAction {
	s.Path = &v
	return s
}

func (s *HTTPGetAction) SetPort(v int32) *HTTPGetAction {
	s.Port = &v
	return s
}

func (s *HTTPGetAction) SetScheme(v string) *HTTPGetAction {
	s.Scheme = &v
	return s
}

func (s *HTTPGetAction) Validate() error {
	if s.HttpHeaders != nil {
		for _, item := range s.HttpHeaders {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
