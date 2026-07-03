// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAlertResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAlert(v *GetAlertResponseBodyAlert) *GetAlertResponseBody
	GetAlert() *GetAlertResponseBodyAlert
	SetRequestId(v string) *GetAlertResponseBody
	GetRequestId() *string
}

type GetAlertResponseBody struct {
	// If this value is true, the minor engine version is not the latest version.
	//
	// > If the minor engine version of your server is not the latest version, the sampling logs may be inaccurate, which causes inaccurate IP statistics. We recommend that you upgrade the minor engine version at your earliest convenience.
	Alert *GetAlertResponseBodyAlert `json:"Alert,omitempty" xml:"Alert,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAlertResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAlertResponseBody) GoString() string {
	return s.String()
}

func (s *GetAlertResponseBody) GetAlert() *GetAlertResponseBodyAlert {
	return s.Alert
}

func (s *GetAlertResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAlertResponseBody) SetAlert(v *GetAlertResponseBodyAlert) *GetAlertResponseBody {
	s.Alert = v
	return s
}

func (s *GetAlertResponseBody) SetRequestId(v string) *GetAlertResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAlertResponseBody) Validate() error {
	if s.Alert != nil {
		if err := s.Alert.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAlertResponseBodyAlert struct {
	// The alert content.
	//
	// example:
	//
	// {***}
	AlertRecord *string `json:"AlertRecord,omitempty" xml:"AlertRecord,omitempty"`
	// The alert UUID.
	//
	// example:
	//
	// sas_1175554881779200_e59368b3ce1690eae7978f7712a69050
	AlertUuid *string `json:"AlertUuid,omitempty" xml:"AlertUuid,omitempty"`
}

func (s GetAlertResponseBodyAlert) String() string {
	return dara.Prettify(s)
}

func (s GetAlertResponseBodyAlert) GoString() string {
	return s.String()
}

func (s *GetAlertResponseBodyAlert) GetAlertRecord() *string {
	return s.AlertRecord
}

func (s *GetAlertResponseBodyAlert) GetAlertUuid() *string {
	return s.AlertUuid
}

func (s *GetAlertResponseBodyAlert) SetAlertRecord(v string) *GetAlertResponseBodyAlert {
	s.AlertRecord = &v
	return s
}

func (s *GetAlertResponseBodyAlert) SetAlertUuid(v string) *GetAlertResponseBodyAlert {
	s.AlertUuid = &v
	return s
}

func (s *GetAlertResponseBodyAlert) Validate() error {
	return dara.Validate(s)
}
