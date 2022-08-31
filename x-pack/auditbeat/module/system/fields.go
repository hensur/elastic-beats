// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package system

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("auditbeat", "system", asset.ModuleFieldsPri, AssetSystem); err != nil {
		panic(err)
	}
}

// AssetSystem returns asset data.
// This is the base64 encoded zlib format compressed contents of module/system.
func AssetSystem() string {
	return "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"
}
