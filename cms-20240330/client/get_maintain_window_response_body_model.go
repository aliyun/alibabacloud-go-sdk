// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMaintainWindowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaintainWindow(v *MaintainWindowForView) *GetMaintainWindowResponseBody
	GetMaintainWindow() *MaintainWindowForView
	SetRequestId(v string) *GetMaintainWindowResponseBody
	GetRequestId() *string
}

type GetMaintainWindowResponseBody struct {
	MaintainWindow *MaintainWindowForView `json:"maintainWindow,omitempty" xml:"maintainWindow,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 0CEC5375-C554-562B-A65F-9A629907C1F0
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetMaintainWindowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMaintainWindowResponseBody) GoString() string {
	return s.String()
}

func (s *GetMaintainWindowResponseBody) GetMaintainWindow() *MaintainWindowForView {
	return s.MaintainWindow
}

func (s *GetMaintainWindowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMaintainWindowResponseBody) SetMaintainWindow(v *MaintainWindowForView) *GetMaintainWindowResponseBody {
	s.MaintainWindow = v
	return s
}

func (s *GetMaintainWindowResponseBody) SetRequestId(v string) *GetMaintainWindowResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMaintainWindowResponseBody) Validate() error {
	if s.MaintainWindow != nil {
		if err := s.MaintainWindow.Validate(); err != nil {
			return err
		}
	}
	return nil
}
