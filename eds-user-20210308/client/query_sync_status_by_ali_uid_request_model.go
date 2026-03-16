// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySyncStatusByAliUidRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessChannel(v string) *QuerySyncStatusByAliUidRequest
	GetBusinessChannel() *string
}

type QuerySyncStatusByAliUidRequest struct {
	// example:
	//
	// ENTERPRISE
	BusinessChannel *string `json:"BusinessChannel,omitempty" xml:"BusinessChannel,omitempty"`
}

func (s QuerySyncStatusByAliUidRequest) String() string {
	return dara.Prettify(s)
}

func (s QuerySyncStatusByAliUidRequest) GoString() string {
	return s.String()
}

func (s *QuerySyncStatusByAliUidRequest) GetBusinessChannel() *string {
	return s.BusinessChannel
}

func (s *QuerySyncStatusByAliUidRequest) SetBusinessChannel(v string) *QuerySyncStatusByAliUidRequest {
	s.BusinessChannel = &v
	return s
}

func (s *QuerySyncStatusByAliUidRequest) Validate() error {
	return dara.Validate(s)
}
