// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRetcodeAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateRetcodeAppResponseBody
	GetRequestId() *string
	SetRetcodeAppDataBean(v *CreateRetcodeAppResponseBodyRetcodeAppDataBean) *CreateRetcodeAppResponseBody
	GetRetcodeAppDataBean() *CreateRetcodeAppResponseBodyRetcodeAppDataBean
}

type CreateRetcodeAppResponseBody struct {
	RequestId          *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RetcodeAppDataBean *CreateRetcodeAppResponseBodyRetcodeAppDataBean `json:"RetcodeAppDataBean,omitempty" xml:"RetcodeAppDataBean,omitempty" type:"Struct"`
}

func (s CreateRetcodeAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRetcodeAppResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRetcodeAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRetcodeAppResponseBody) GetRetcodeAppDataBean() *CreateRetcodeAppResponseBodyRetcodeAppDataBean {
	return s.RetcodeAppDataBean
}

func (s *CreateRetcodeAppResponseBody) SetRequestId(v string) *CreateRetcodeAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRetcodeAppResponseBody) SetRetcodeAppDataBean(v *CreateRetcodeAppResponseBodyRetcodeAppDataBean) *CreateRetcodeAppResponseBody {
	s.RetcodeAppDataBean = v
	return s
}

func (s *CreateRetcodeAppResponseBody) Validate() error {
	if s.RetcodeAppDataBean != nil {
		if err := s.RetcodeAppDataBean.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateRetcodeAppResponseBodyRetcodeAppDataBean struct {
	AppId *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Pid   *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
}

func (s CreateRetcodeAppResponseBodyRetcodeAppDataBean) String() string {
	return dara.Prettify(s)
}

func (s CreateRetcodeAppResponseBodyRetcodeAppDataBean) GoString() string {
	return s.String()
}

func (s *CreateRetcodeAppResponseBodyRetcodeAppDataBean) GetAppId() *int64 {
	return s.AppId
}

func (s *CreateRetcodeAppResponseBodyRetcodeAppDataBean) GetPid() *string {
	return s.Pid
}

func (s *CreateRetcodeAppResponseBodyRetcodeAppDataBean) SetAppId(v int64) *CreateRetcodeAppResponseBodyRetcodeAppDataBean {
	s.AppId = &v
	return s
}

func (s *CreateRetcodeAppResponseBodyRetcodeAppDataBean) SetPid(v string) *CreateRetcodeAppResponseBodyRetcodeAppDataBean {
	s.Pid = &v
	return s
}

func (s *CreateRetcodeAppResponseBodyRetcodeAppDataBean) Validate() error {
	return dara.Validate(s)
}
