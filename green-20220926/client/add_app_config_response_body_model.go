// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAppConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *AddAppConfigResponseBodyData) *AddAppConfigResponseBody
	GetData() *AddAppConfigResponseBodyData
	SetRequestId(v string) *AddAppConfigResponseBody
	GetRequestId() *string
}

type AddAppConfigResponseBody struct {
	// The returned data.
	Data *AddAppConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID assigned by the backend to uniquely identify a request. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddAppConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddAppConfigResponseBody) GoString() string {
	return s.String()
}

func (s *AddAppConfigResponseBody) GetData() *AddAppConfigResponseBodyData {
	return s.Data
}

func (s *AddAppConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddAppConfigResponseBody) SetData(v *AddAppConfigResponseBodyData) *AddAppConfigResponseBody {
	s.Data = v
	return s
}

func (s *AddAppConfigResponseBody) SetRequestId(v string) *AddAppConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddAppConfigResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddAppConfigResponseBodyData struct {
	// App ID。
	//
	// example:
	//
	// txt_check_pro_agent_01
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s AddAppConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AddAppConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddAppConfigResponseBodyData) GetAppId() *string {
	return s.AppId
}

func (s *AddAppConfigResponseBodyData) SetAppId(v string) *AddAppConfigResponseBodyData {
	s.AppId = &v
	return s
}

func (s *AddAppConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}
