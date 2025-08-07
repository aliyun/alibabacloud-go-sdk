// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterServerlessConfResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ModifyDBClusterServerlessConfResponseBody
	GetDBClusterId() *string
	SetRequestId(v string) *ModifyDBClusterServerlessConfResponseBody
	GetRequestId() *string
}

type ModifyDBClusterServerlessConfResponseBody struct {
	// The ID of the serverless cluster.
	//
	// example:
	//
	// pc-bp10gr51qasnl****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 5E71541A-6007-4DCC-A38A-F872C31FEB45
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBClusterServerlessConfResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterServerlessConfResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterServerlessConfResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyDBClusterServerlessConfResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBClusterServerlessConfResponseBody) SetDBClusterId(v string) *ModifyDBClusterServerlessConfResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterServerlessConfResponseBody) SetRequestId(v string) *ModifyDBClusterServerlessConfResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBClusterServerlessConfResponseBody) Validate() error {
	return dara.Validate(s)
}
