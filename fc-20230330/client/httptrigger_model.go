// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHTTPTrigger interface {
	dara.Model
	String() string
	GoString() string
	SetUrlInternet(v string) *HTTPTrigger
	GetUrlInternet() *string
	SetUrlIntranet(v string) *HTTPTrigger
	GetUrlIntranet() *string
}

type HTTPTrigger struct {
	// example:
	//
	// https://svc-func-xxxxxxxx.cn-hangzhou.fcapp.run
	UrlInternet *string `json:"urlInternet,omitempty" xml:"urlInternet,omitempty"`
	// example:
	//
	// https://svc-func-xxxxxxxx.cn-hangzhou-vpc.fcapp.run
	UrlIntranet *string `json:"urlIntranet,omitempty" xml:"urlIntranet,omitempty"`
}

func (s HTTPTrigger) String() string {
	return dara.Prettify(s)
}

func (s HTTPTrigger) GoString() string {
	return s.String()
}

func (s *HTTPTrigger) GetUrlInternet() *string {
	return s.UrlInternet
}

func (s *HTTPTrigger) GetUrlIntranet() *string {
	return s.UrlIntranet
}

func (s *HTTPTrigger) SetUrlInternet(v string) *HTTPTrigger {
	s.UrlInternet = &v
	return s
}

func (s *HTTPTrigger) SetUrlIntranet(v string) *HTTPTrigger {
	s.UrlIntranet = &v
	return s
}

func (s *HTTPTrigger) Validate() error {
	return dara.Validate(s)
}
