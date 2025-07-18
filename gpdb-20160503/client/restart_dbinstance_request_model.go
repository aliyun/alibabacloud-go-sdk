// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartDBInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *RestartDBInstanceRequest
	GetClientToken() *string
	SetDBInstanceId(v string) *RestartDBInstanceRequest
	GetDBInstanceId() *string
}

type RestartDBInstanceRequest struct {
	// The client token that is used to ensure the idempotence of the request. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/327176.html).
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88**********
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the instance.
	//
	// >  You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the details of all AnalyticDB for PostgreSQL instances in a specific region, including instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s RestartDBInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RestartDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *RestartDBInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RestartDBInstanceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *RestartDBInstanceRequest) SetClientToken(v string) *RestartDBInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *RestartDBInstanceRequest) SetDBInstanceId(v string) *RestartDBInstanceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *RestartDBInstanceRequest) Validate() error {
	return dara.Validate(s)
}
