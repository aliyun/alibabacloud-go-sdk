// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncComponentTemplateConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SyncComponentTemplateConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SyncComponentTemplateConfigResponseBody
	GetSuccess() *bool
}

type SyncComponentTemplateConfigResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 3239281273464326823
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// ture
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SyncComponentTemplateConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SyncComponentTemplateConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SyncComponentTemplateConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SyncComponentTemplateConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SyncComponentTemplateConfigResponseBody) SetRequestId(v string) *SyncComponentTemplateConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *SyncComponentTemplateConfigResponseBody) SetSuccess(v bool) *SyncComponentTemplateConfigResponseBody {
	s.Success = &v
	return s
}

func (s *SyncComponentTemplateConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
