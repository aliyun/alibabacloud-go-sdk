// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAlertMessageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertId(v string) *GetAlertMessageRequest
	GetAlertId() *string
}

type GetAlertMessageRequest struct {
	// The alert ID. You can all the [ListAlertMessages](https://help.aliyun.com/document_detail/173961.html) operation to obtain the alert ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1421
	AlertId *string `json:"AlertId,omitempty" xml:"AlertId,omitempty"`
}

func (s GetAlertMessageRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAlertMessageRequest) GoString() string {
	return s.String()
}

func (s *GetAlertMessageRequest) GetAlertId() *string {
	return s.AlertId
}

func (s *GetAlertMessageRequest) SetAlertId(v string) *GetAlertMessageRequest {
	s.AlertId = &v
	return s
}

func (s *GetAlertMessageRequest) Validate() error {
	return dara.Validate(s)
}
