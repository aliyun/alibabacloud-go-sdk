// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebalanceDBInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *RebalanceDBInstanceRequest
	GetClientToken() *string
	SetDBInstanceId(v string) *RebalanceDBInstanceRequest
	GetDBInstanceId() *string
}

type RebalanceDBInstanceRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that it is unique among different requests.
	//
	// The token can be up to 64 characters in length and can contain letters, digits, hyphens (-), and underscores (_).
	//
	// For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/134212.html).
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88**********
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-bp***************
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s RebalanceDBInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RebalanceDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *RebalanceDBInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RebalanceDBInstanceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *RebalanceDBInstanceRequest) SetClientToken(v string) *RebalanceDBInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *RebalanceDBInstanceRequest) SetDBInstanceId(v string) *RebalanceDBInstanceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *RebalanceDBInstanceRequest) Validate() error {
	return dara.Validate(s)
}
